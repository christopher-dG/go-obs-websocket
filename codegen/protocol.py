# RUN GOFMT BEFORE LOOKING AT THE OUTPUT PLEASE.

import json
import os
import sys
from typing import Dict, List, Tuple

package = "obsws"

type_map = {
    "boolean": "bool",
    "int": "int",
    "double": "float64",
    "string": "string",
    "array": "[]string",
    "object": "map[string]interface{}",
    "array of objects": "[]map[string]interface{}",
    "scene": "interface{}",  # String?
    "object|array": "interface{}",
    "scene|array": "interface{}",
    "source|array": "interface{}",
}

unknown_types = [
    "object|array",
    "scene|array",
    "scene",  # String?
    "source|array",
]


def optional_type(s: str) -> Tuple[str, bool]:
    if s.endswith("(optional)"):
        return s[:s.find("(optional)")].strip(), True
    return s, False


def process_json(d: Dict):
    process_events(d["events"])
    process_requests(d["requests"])


def process_events(events: Dict):
    for category, data in events.items():
        process_events_category(category, data)


def process_events_category(category: str, data: Dict):
    events = "\n\n".join(generate_event(event) for event in data)
    with open(go_filename("events", category), "w") as f:
        f.write(f"""\
        package {package}

        // This code is automatically generated.
        // See: https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

        // https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#{category}

        {events}
        """)


def generate_event(data: Dict) -> str:
    """Generate Go code with type definitions and interface functions."""
    if "returns" in data:
        struct = f"""\
        type {data['name']}Event struct {{
            {go_variables(data['returns'])}
            _event
        }}\
        """
    else:
        struct = f"type {data['name']}Event _event"

    description = data["description"].replace("\n", " ")
    description = f"{data['name']}Event : {description}"
    if "since" in data:
        description += f" Since: {data['since'].capitalize()}"

    return f"""\
    // {description}
    // https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#{data['heading']['text'].lower()}
    {struct}

    // Type returns the event's update type.
    func (e {data['name']}Event) Type() string {{ return e.UpdateType }}

    // StreamTC returns the event's stream timecode.
    func (e {data['name']}Event) StreamTC() string {{ return e.StreamTimecode }}

    // RecTC returns the event's recording timecode.
    func (e {data['name']}Event) RecTC() string {{ return e.RecTimecode }}
    """


def process_requests(requests: Dict):
    pass


def go_variables(vars: List) -> str:
    """
    Convert a list of variable definition into Go code to be put
    inside a struct definition.
    """
    lines = []
    for v in vars:
        line = go_name(v["name"])
        typename, optional = optional_type(v["type"])
        line += f" {type_map[typename.lower()]} // {v['description']}"
        if optional:
            line += " Optional."
        if typename.lower() in unknown_types:
            line += f" TODO: Unknown type ({typename})."
        lines.append(line)
    return "\n".join(lines)


def go_name(s: str) -> str:
    """
    Convert a variable name in the input file to a Go variable name.
    Note: This makes lots of assumptions about the input,
    i.e. nothing ends with a separator.
    """
    s = s.capitalize()
    for sep in ["-", "_", ".*.", "[].", "."]:
        while sep in s:
            i = s.find(sep)
            _len = len(sep)
            s = f"{s[:i]}{s[i+_len].upper()}{s[i+_len+1:]}"
    return s


def go_filename(category, section):
    return f"{category}_{section.replace(' ', '_')}.go"


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Missing filename argument")
        exit(1)

    if not os.path.isfile(sys.argv[1]):
        print(f"file '{sys.argv[1]}' does not exist")
        exit(1)

    with open(sys.argv[1]) as f:
        d = json.load(f)

    process_json(d)
