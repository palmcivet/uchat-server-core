# /usr/bin/env python3
# codeing=utf-8

import os


def print_promt(type, prompt):
    prompt = ""
    return prompt


def arg_query():
    arg = {
        "author": "Palm Civet",
        "pwd": "./",
        "project name": "test",
        "type": "react",
    }
    for (key, value) in arg.items():
        try:
            value = input("\033[1;32;40m%s(%s): \033[0m" % (key, value))
        except IOError:
            value = value
    return arg


def main():
    arg = arg_query()
    for key, value in arg.items():
        print(key + ": " + value)

    try:
        os.makedirs(os.path.join(arg["pwd"], arg["project name"]))
    except FileExistsError:
        print("\033[32;0mCan't make a folder, please check the dir\033[0m")

    for files in os.listdir(arg["pwd"]):
        if os.path.isdir(files):
            print(files)


if __name__ == "__main__":
    main()
