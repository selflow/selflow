"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[26],{3905:(e,t,n)=>{n.d(t,{Zo:()=>m,kt:()=>f});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},l=Object.keys(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),d=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},m=function(e){var t=d(e.components);return a.createElement(s.Provider,{value:t},e.children)},u="mdxType",p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},c=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,l=e.originalType,s=e.parentName,m=i(e,["components","mdxType","originalType","parentName"]),u=d(n),c=r,f=u["".concat(s,".").concat(c)]||u[c]||p[c]||l;return n?a.createElement(f,o(o({ref:t},m),{},{components:n})):a.createElement(f,o({ref:t},m))}));function f(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=n.length,o=new Array(l);o[0]=c;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[u]="string"==typeof e?e:r,o[1]=i;for(var d=2;d<l;d++)o[d]=n[d];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}c.displayName="MDXCreateElement"},5026:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>m,contentTitle:()=>s,default:()=>c,frontMatter:()=>i,metadata:()=>d,toc:()=>u});var a=n(3117),r=n(102),l=(n(7294),n(3905)),o=["components"],i={title:"\ud83d\udc31 Selflow-Daemon"},s=void 0,d={unversionedId:"ecosystem/selflow-daemon",id:"ecosystem/selflow-daemon",title:"\ud83d\udc31 Selflow-Daemon",description:"Added features",source:"@site/../../docs/ecosystem/selflow-daemon.mdx",sourceDirName:"ecosystem",slug:"/ecosystem/selflow-daemon",permalink:"/selflow/docs/ecosystem/selflow-daemon",draft:!1,editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/../../docs/ecosystem/selflow-daemon.mdx",tags:[],version:"current",frontMatter:{title:"\ud83d\udc31 Selflow-Daemon"},sidebar:"tutorialSidebar",previous:{title:"\ud83d\udd79\ufe0f CLI",permalink:"/selflow/docs/ecosystem/cli"},next:{title:"\ud83c\udf10 Webclient",permalink:"/selflow/docs/ecosystem/webclient"}},m={},u=[{value:"Added features",id:"added-features",level:2},{value:"Requirements",id:"requirements",level:2},{value:"Getting started",id:"getting-started",level:2},{value:"Customisation",id:"customisation",level:2}],p={toc:u};function c(e){var t=e.components,n=(0,r.Z)(e,o);return(0,l.kt)("wrapper",(0,a.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h2",{id:"added-features"},"Added features"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"GRPC API"),(0,l.kt)("li",{parentName:"ul"},"Support for Docker Steps")),(0,l.kt)("p",null,"Selflow-Daemon is, as of today, the primary implementation of Selflow.\nIt adds capabilities for docker steps and can communicate with a GRPC API."),(0,l.kt)("p",null,"It runs on port ",(0,l.kt)("inlineCode",{parentName:"p"},"10011")," in a docker container to allows you to remove it easily.\nIt still needs to access a folder on your machine.\nUsually the ",(0,l.kt)("inlineCode",{parentName:"p"},"/etc/selflow")," folder but you can customize it using environment variables"),(0,l.kt)("h2",{id:"requirements"},"Requirements"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"a machine running Linux or MacOS (It works on windows but is not stable enough for now)"),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"https://docs.docker.com/engine/install/"},"docker installed and setup"))),(0,l.kt)("h2",{id:"getting-started"},"Getting started"),(0,l.kt)("p",null,"The easiest way to start the Selflow-Daemon is to use the ",(0,l.kt)("a",{parentName:"p",href:"./cli"},"CLI")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"selflow recreate-daemon\n")),(0,l.kt)("p",null,"The Deamon will be running on port 10011 on your system and run logs will be stored in the ",(0,l.kt)("inlineCode",{parentName:"p"},"/etc/selflow")," directory on your system."),(0,l.kt)("h2",{id:"customisation"},"Customisation"),(0,l.kt)("p",null,"You can customize the Selflow Daemon behavior using environment variables that you can directly add to the Selflow CLI"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Variable"),(0,l.kt)("th",{parentName:"tr",align:null},"Default Value"),(0,l.kt)("th",{parentName:"tr",align:null},"Role"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"SELFLOW_DAEMON_PORT")),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"10011")),(0,l.kt)("td",{parentName:"tr",align:null},"Port where the Selflow Daemon will run")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"SELFLOW_DAEMON_NAME")),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"selflow-daemon")),(0,l.kt)("td",{parentName:"tr",align:null},"Name of the container")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"SELFLOW_DAEMON_NETWORK")),(0,l.kt)("td",{parentName:"tr",align:null},"Same as ",(0,l.kt)("inlineCode",{parentName:"td"},"SELFLOW_DAEMON_NAME")),(0,l.kt)("td",{parentName:"tr",align:null},"Name of the network")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"SELFLOW_DAEMON_IMAGE")),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"selflow-daemon:latest")),(0,l.kt)("td",{parentName:"tr",align:null},"Name of the docker image")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"SELFLOW_DAEMON_HOST_BASED_DIRECTORY")),(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"/etc/selflow")),(0,l.kt)("td",{parentName:"tr",align:null},"Name of the directory on host where files will be stored. It will be mapped as a volume")))))}c.isMDXComponent=!0}}]);