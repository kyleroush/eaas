(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{1:function(e,t,a){e.exports=a.p+"static/media/logo.5d5d9eef.svg"},15:function(e,t,a){},16:function(e,t,a){},17:function(e,t,a){"use strict";a.r(t);var n=a(0),o=a.n(n),r=a(3),s=a.n(r),c=(a(15),a(4)),l=a(5),i=a(7),p=a(6),m=a(8),d=a(1),u=a.n(d),h=(a(16),function(e){function t(e){var a;return Object(c.a)(this,t),(a=Object(i.a)(this,Object(p.a)(t).call(this,e))).state={error:null,isLoaded:!1,result:{}},a}return Object(m.a)(t,e),Object(l.a)(t,[{key:"componentDidMount",value:function(){var e=this;fetch("https://7siim5bbqh.execute-api.us-east-1.amazonaws.com/Prod?format=list",{mode:"no-cors",headers:{"Content-Type":"application/json"}}).then(function(e){return e.json()}).then(function(t){e.setState({isLoaded:!0,result:t})},function(t){e.setState({isLoaded:!1,error:t.message})})}},{key:"render",value:function(){return this.state.isLoaded?(console.log(this.state.result),o.a.createElement("div",{className:"App"},o.a.createElement("header",{className:"App-header"},o.a.createElement("img",{src:u.a,className:"App-logo",alt:"logo"}),o.a.createElement("p",null,"HELLO"),o.a.createElement("a",{className:"App-link",href:"https://reactjs.org",target:"_blank",rel:"noopener noreferrer"})))):o.a.createElement("div",{className:"App"},o.a.createElement("header",{className:"App-header"},o.a.createElement("img",{src:u.a,className:"App-logo",alt:"logo"}),o.a.createElement("p",null,"Edit ",o.a.createElement("code",null,"src/App.js")," and save to reload."),o.a.createElement("a",{className:"App-link",href:"https://reactjs.org",target:"_blank",rel:"noopener noreferrer"},"Learn React")))}}]),t}(n.Component));Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));s.a.render(o.a.createElement(h,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})},9:function(e,t,a){e.exports=a(17)}},[[9,1,2]]]);
//# sourceMappingURL=main.33330a2e.chunk.js.map