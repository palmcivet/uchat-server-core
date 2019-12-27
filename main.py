# /usr/bin/env python3
# codeing=utf-8

import os
import random  # TODO


# 打印信息：错误、警告、正常
def print_prompt(prompt_type):
    if prompt_type == "Query":
        rtn_prompt = config_type["Begin"] + config_type["Query"]
        rtn_prompt += "%s("
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type["Quote"]
        rtn_prompt += "%s"
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type["Query"]
        rtn_prompt += "):"
        rtn_prompt += config_type["End"]
    elif prompt_type == "Normal":
        rtn_prompt = config_type["Begin"] + config_type["Query"]
        rtn_prompt += "%s("
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type["Quote"]
        rtn_prompt += "%s"
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type["Query"]
        rtn_prompt += "):"
        rtn_prompt += config_type["End"]
    elif prompt_type == "Error":
        rtn_prompt = config_type["Begin"] + config_type["Query"]
        rtn_prompt += "["
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type[prompt_type]
        rtn_prompt += prompt_type
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type["Query"]
        rtn_prompt += "] "
        rtn_prompt += config_type["End"]
        rtn_prompt += config_type["Begin"] + config_type["Quote"]
        rtn_prompt += "%s"
        rtn_prompt += config_type["End"]
    else:
        rtn_prompt = config_type["Begin"] + config_type["Normal"]
    return rtn_prompt


# 选择不同项目，目前只有 front_end
def operate(new_type):
    for (item_key, item_val) in front_end.items():
        if new_type in item_val.keys():
            os.system("yarn init")
            # print("init completed.")  # TODO
            for lib in item_val[new_type]:
                try:
                    os.system("yarn add %s" % lib)
                    # print("yarn add %s" % lib)  # TODO
                except IOError:
                    print(
                        print_prompt("Error") %
                        "Stopped at %s. Check and go on." % lib)
                    exit(0)
            continue

        a = input(
            print_prompt("Query") % ("Select an item", "|".join(
                i.replace(" -D", "") for i in item_val)))
        for lib in item_val[a]:
            try:
                # print("yarn add %s" % lib)  # TODO
                os.system("yarn add %s" % lib)
            except IOError:
                print(
                    print_prompt("Error") % "Stopped at %s. Check and go on." %
                    lib)
                exit(0)
    print_prompt("Settle down. Hack fun!")


# 初始化环境
def env_init():
    arg = {
        # "author": "Palm Civet",
    }
    for (key, value) in arg.items():
        try:
            value = input(print_prompt("Query") % (key, value))
        except IOError:
            value = value
    return arg


# 项目类型、项目文件夹路径
def arg_query():
    arg = {
        "type": "react",
        # "name": "./test",
        "name": "./test/" + str(random.randrange(1, 50)),  # TODO
    }
    for (key, value) in arg.items():
        arg[key] = input(print_prompt("Query") % (key, value)) or value
    return arg


def main():
    arg = arg_query()
    try:
        os.makedirs(os.path.join("./", arg["name"].replace("./", "")))
        os.chdir(os.path.join("./", arg["name"].replace("./", "")))
    except FileExistsError:
        print(print_prompt("Error") % "Can't make a folder.")
        exit(0)
    operate(arg["type"])

    for files in os.listdir():
        if os.path.isdir(files):
            print(files)


if __name__ == "__main__":
    config_type = {
        "Begin": "\033[1;",
        "End": "\033[0m",
        "Error": "31;40m",
        "Normal": "32;40m",
        "Warn": "33;40m",
        "Query": "37;40m",
        "Quote": "34;40m",
    }
    front_end = {
        "library": {
            "react": [
                "react",
                "react-dom",
                "@babel/cli -D",
                "@babel/core -D",
                "@babel/preset-env -D",
                "@babel/preset-react -D",
                "babel-loader -D",
                "css-loader -D",
                "style-loader -D",
                "webpack -D",
                "webpack-cli -D",
                "webpack-dev-server -D",
                "html-webpack-plugin -D",
            ],
            "vue": [],
        },
        "css": {
            "less": ["less -D", "less-loader -D"],
            "sass": ["sass -D", "sass-loader -D"],
        },
        "ui": {
            "antd":
            ["antd -D", "react-hot-loader -D", "babel-plugin-import -D"],
            "element": ["element-react -D", "element-theme-default -D"],
        },
        "data": {
            "mobx": [
                "babel-plugin-transform-class-properties -D",
                "babel-plugin-transform-decorators-legacy -D",
                "@babel/plugin-proposal-decorators -D",
            ],
            "redux": ["redux -D", "react-redux -D"],
        }
    }
    main()
