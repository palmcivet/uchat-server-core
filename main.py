# /usr/bin/env python3
# codeing=utf-8
# author: Palm Civet
# repo: https://www.github.com/Palmcivet.git

import os
import random  # TODO


class gen_prompt(object):
    __ctrl = {
        "Begin": "\033[1;",
        "End": "\033[0m",
        "Quote": "34;40m",
        "Body": "37;40m",
    }

    __sym = {
        "Success": "36;40m" + "√",  # cyan
        "Query": "32;40m" + "?",  # green
        "Error": "31;40m" + "×",  # red
        "Warn": "33;40m" + "!",  # yellow
    }

    # type - Warn|Error|Query|Success
    def __prompt_sym(self, sym_type):
        sym = self.__ctrl["Begin"] + self.__ctrl["Body"] + "[" + self.__ctrl[
            "End"]
        sym += self.__ctrl["Begin"] + self.__sym[sym_type] + self.__ctrl["End"]
        sym += self.__ctrl["Begin"] + self.__ctrl["Body"] + "]" + self.__ctrl[
            "End"]
        sym += " "
        return sym

    # [√] Settle down. hack fun!
    def success(self):
        body = self.__prompt_sym("Success")
        body += self.__ctrl["Begin"] + self.__ctrl["Body"]
        body += "%s"
        body += self.__ctrl["End"]
        return body

    # [?] Select lib(less|sass):
    def query(self):
        body = self.__prompt_sym("Query")
        body += self.__ctrl["Begin"] + self.__ctrl["Body"]
        body += "%s("
        body += self.__ctrl["End"]
        body += self.__ctrl["Begin"] + self.__ctrl["Quote"]
        body += "%s"
        body += self.__ctrl["End"]
        body += self.__ctrl["Begin"] + self.__ctrl["Body"]
        body += "):"
        body += self.__ctrl["End"]
        return body

    # [×] Network fail. Stopped at __
    def error(self):
        body = self.__prompt_sym("Error")
        body += self.__ctrl["Begin"] + self.__ctrl["Body"]
        body += "%s"
        body += self.__ctrl["End"]
        body += self.__ctrl["Begin"] + self.__ctrl["Quote"]
        body += "%s"
        body += self.__ctrl["End"]
        return body

    def warn(self):
        body = self.__prompt_sym("Warn")
        body += self.__ctrl["Begin"] + self.__ctrl["Body"]
        body += "%s"
        body += self.__ctrl["End"]
        return body


class init_scaffold(object):
    def __init__(self):
        pass

    def env_query(self, hdl_notice):
        env = {
            "git": "git",
            "npm": "npm",
            "yarn": "yarn",
            "author": "Palm Civet",
        }
        return env

    def arg_query(self, hdl_notice):
        arg = {
            "type": "react",
            # "name": "./test",
            "name": "./test/" + str(random.randrange(1, 50)),  # TODO
        }
        for (key, value) in arg.items():
            arg[key] = input(
                (hdl_notice.query() + " ") % (key, value)) or value
        return arg

    def git_init(self):
        # os.system(env["git"])
        print("git init")  # TODO


class op_directory(object):
    def __init__(self):
        pass

    def list_dir(self, dir="./", op="l", hide=False):
        if op == "l":
            for files in os.listdir():
                if os.path.isdir(files):
                    print(files)
        elif op == "t":
            pass
        else:
            pass

    def mk_cd(self, file_name):
        try:
            os.makedirs(os.path.join("./", file_name.replace("./", "")))
            os.chdir(os.path.join("./", file_name.replace("./", "")))
        except FileExistsError:
            print("Can't make a folder.")
            exit(0)


class operate(object):
    __front_end = {
        "Library": {
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
            "vue": [
                "vue",
            ],
        },
        "CSS": {
            "less": ["less -D", "less-loader -D"],
            "sass": ["sass -D", "sass-loader -D"],
        },
        "UI": {
            "antd":
            ["antd -D", "react-hot-loader -D", "babel-plugin-import -D"],
            "element": ["element-react -D", "element-theme-default -D"],
        },
        "Data": {
            "mobx": [
                "babel-plugin-transform-class-properties -D",
                "babel-plugin-transform-decorators-legacy -D",
                "@babel/plugin-proposal-decorators -D",
            ],
            "redux": ["redux -D", "react-redux -D"],
        }
    }

    def front_end(self, prompt, arg, env):
        for (item_key, item_val) in self.__front_end.items():
            scheme = arg["type"]
            if scheme in item_val.keys():
                # os.system("%s init" % env["yarn"])
                print("yarn init.")  # TODO
            else:
                scheme = input((prompt.query() + " ") % (item_key, "|".join(
                    i.replace(" -D", "") for i in item_val)))

                while scheme not in item_val:
                    if scheme in ["q", "e", "quit", "exit"]:
                        exit(0)
                    else:
                        print(prompt.error() % ("Don't exist: ", scheme))
                    scheme = input(
                        prompt.warn() %
                        ("Choose an item(\033[1;34;40m%s\033[0m): ") %
                        "|".join(i.replace(" -D", "") for i in item_val))

            for lib in item_val[scheme]:
                try:
                    # os.system("yarn add %s" % lib)
                    print("yarn add %s" % lib)  # TODO
                except IOError:
                    print(prompt.error() % ("Stopped at: ", lib))
                    exit(0)

    def python(self, parameter_list):
        pass

    def rust(self, parameter_list):
        pass


def main():
    notice = gen_prompt()

    scaffold = init_scaffold()
    env = scaffold.env_query(notice)  # config environment
    arg = scaffold.arg_query(notice)  # query arguments

    op = op_directory()
    op.mk_cd(arg["name"])  # mkdir and cd
    scaffold.git_init()  # git init

    hdl = operate()
    hdl.front_end(notice, arg, env)

    print(notice.success() % "Settle down. Hack fun!")


if __name__ == "__main__":
    main()
