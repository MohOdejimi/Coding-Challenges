from pathlib import Path 
import os 
import sys 


def safeToProcess(filename):
    canProcess = True
    msg = ""

    if not filename.is_file():
        canProcess = False
        msg = f"Error: {filename} is not a valid file"
    elif os.stat(filename).st_size == 0:
        canProcess = False 
        msg = f"Error: {filename} is empty"
    elif not os.access(filename, os.F_OK):
        canProcess = False
        msg = f"Error: {filename} doesn't exist"

    return (canProcess, msg) 

def processFile(filename):
    content = ""
    with open(filename) as file:
        content += file.read()
    return content
        
def main():
    args = sys.argv[1:]
    content = ""
    failedToProcess =  False
    errorMsg = ""
    is_standardinput_empty = sys.stdin.isatty()

    if is_standardinput_empty:
        for file in args:
            filename = Path(file)
            canProcess, msg = safeToProcess(filename)
            if not canProcess:
                failedToProcess = True 
                errorMsg = msg
            else:
                content += processFile(filename)
    else:
        standardInput = sys.stdin.read().strip().split("\n")
        num = 0
        for line in standardInput:
            print(f"{num + 1}. {line}")
            num += 1

    if failedToProcess:
        print(errorMsg)
        if content != "":
            print(content, end = "")
        sys.exit(1)

    else:
        print(content, end="")


main()