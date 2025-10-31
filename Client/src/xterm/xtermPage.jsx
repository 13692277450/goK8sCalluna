import React, { Component } from "react";
import { Terminal } from "xterm";
import "xterm/css/xterm.css";
import { getToken } from "@libs/auth";
import { FitAddon } from "xterm-addon-fit";

export default class WebTerminal extends Component {
  term = null;
  websocket = null;
  curr_line = "";
  componentDidMount() {
    let term = this.term;
    this.term = new Terminal({
      fontFamily: 'Menlo, Monaco, "Courier New", monospace',
      fontWeight: 400,
      fontSize: 14,
      rows: Math.ceil(
        (document.getElementsByClassName("container-children")[0].clientHeight -
          150) /
          14
      ),
    });
    this.term.open(document.getElementById("terminal"));
    this.term.focus();
    this.term.prompt = (_) => {
      this.term.write("\r\n\x1b[33m$\x1b[0m ");
    };
    this.term.prompt();
    const fitAddon = new FitAddon();
    this.term.loadAddon(fitAddon);
    fitAddon.fit();
    this.term.prompt();
    this.term.attachCustomKeyEventHandler((e) => {
      console.log({ e });
      //   e = e.target;
      var keyCode = e.keyCode || e.which || e.charCode;
      const moveKey = [37, 38, 39, 40].includes(keyCode);
      if (moveKey) return false;
    });
    // 添加事件监听器，支持输入方法
    this.term.onKey((e) => {
      const printable =
        !e.domEvent.altKey &&
        !e.domEvent.altGraphKey &&
        !e.domEvent.ctrlKey &&
        !e.domEvent.metaKey;
      if (e.domEvent.keyCode === 13) {
        this.Send(term, this.curr_line);
        this.term.prompt();
        this.curr_line = "";
      } else if (e.domEvent.keyCode === 8) {
        // back 删除的情况
        if (this.term._core.buffer.x > 2) {
          if (this.curr_line.length) {
            this.curr_line = this.curr_line.slice(0, this.curr_line.length - 1);
            this.term.write("\b \b");
          } else {
          }
        }
      } else if (printable) {
        this.curr_line += e.key;
        this.term.write(e.key);
      }
      this.term.focus();
      console.log(1, "print", e.key);
    });
    this.term.onData((key) => {
      // 粘贴的情况
      if (key.length > 1) {
        this.term.write(key);
        this.curr_line += key;
      }
    });
    this.initWebsock();
  }
  componentWillUnmount() {
    this.term.dispose();
    this.websocket.close()
  }
  initWebsock = () => {
    // let websocket = this.websocket;
    let term = this.term;
    let token = getToken();
    this.websocket = new WebSocket("ws://" + window._CONFIG.WsSsh, token);
    this.websocket.onopen = function (evt) {
      term.write("connect");
    };
    this.websocket.onclose = function (evt) {
      term.write("exit");
    };
    this.websocket.onmessage = function (evt) {
      term.write(evt.data);
    };
    this.websocket.onerror = function (evt) {
      term.write("connect fail err:" + evt.data);
    };
  };
  //   prompt = (term) => {
  //     this.term.write("\r\n~$ ");
  //   };

  Send = (term, message) => {
    this.websocket.send(message);
  };
  render() {
    return (
      <div className="container-children" style={{ height: "100%" }}>
        <div id="terminal" style={{ width: "100%" }}></div>
      </div>
    );
  }
}
