import argparse 
from pathlib import Path 
import sys
import os 

def safeToProcess(filename):
    canProcess = True 
    msg = ""
    if not filename.is_file():
        canProcess = False 
        msg = f"Error: {filename}is not a file"
    elif os.stat(filename).st_size == 0:
        canProcess = False 
        msg = f"Error: {filename} is currently empty"
    elif not os.access(filename, os.F_OK):
        canProcess = False
        msg = f"Error: {filename} doe's not exist"

    return (canProcess, msg)

def parser():
    parser = argparse.ArgumentParser()

    parser.add_argument('-f')
    parser.add_argument('-d')
    parser.add_argument('filename')

    args = parser.parse_args()

    file = args.filename 
    filename = Path(file)
    fields = args.f
    delimiter = args.d

    return (file, filename, fields, delimiter)

def printContent(input, delimiter, fields):
        for line in input:
            values = line.strip().split(f"{delimiter}")
            fields_entry = []
            for field in fields:
                if field.isdigit():
                    field = int(field)
                    fields_entry.append(values[field - 1])
            print(f'{delimiter}'.join(fields_entry))

def main():
    file, filename, fields, delimiter = parser()

    if ',' in fields:
        fields = fields.split(',')
    else:
        fields = fields.split(' ')

    if not delimiter: 
        delimiter = '\t'

    is_standardinput_empty = sys.stdin.isatty()

    if is_standardinput_empty:
        canProcess, msg = safeToProcess(filename)

        if canProcess:
            with open(filename) as file:
                printContent(file, delimiter, fields)
    else:
        standardInput = sys.stdin.read().strip().split('\n')
        printContent(standardInput, delimiter, fields)
    
main()