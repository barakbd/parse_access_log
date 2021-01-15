# Parse log file to count occurence of status_code
# Usage: python parse_log.py access.log
# See pretty print options below

from sys import argv
from re import split
# Using dumps to ensure dict output with double quotes - json standard
from json import dumps


def parse_logs():
    in_file = argv[1]
    status_code_count_dict = {}
    with open(in_file, "r") as input_file:
        for line in input_file.readlines():
            status_code = str(split(r"\t", line)[5])
            if status_code in status_code_count_dict:
                status_code_count_dict[status_code] += 1
            else:
                status_code_count_dict[status_code] = 1
    return status_code_count_dict


if __name__ == "__main__":
    status_code_count_dict = parse_logs()
    # Pretty print in shell using jq:
    # python parse_log.py access.log | jq .
    print(dumps(status_code_count_dict))
    
    # Pretty print in python (Can also use pprint.PrettyPrinter)
    # print(dumps(status_code_count_dict, indent=2))

