"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[917],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>f});var o=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},l=Object.keys(e);for(o=0;o<l.length;o++)n=l[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(o=0;o<l.length;o++)n=l[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=o.createContext({}),c=function(e){var t=o.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},u=function(e){var t=c(e.components);return o.createElement(s.Provider,{value:t},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},m=o.forwardRef((function(e,t){var n=e.components,r=e.mdxType,l=e.originalType,s=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),p=c(n),m=r,f=p["".concat(s,".").concat(m)]||p[m]||d[m]||l;return n?o.createElement(f,a(a({ref:t},u),{},{components:n})):o.createElement(f,a({ref:t},u))}));function f(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=n.length,a=new Array(l);a[0]=m;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[p]="string"==typeof e?e:r,a[1]=i;for(var c=2;c<l;c++)a[c]=n[c];return o.createElement.apply(null,a)}return o.createElement.apply(null,n)}m.displayName="MDXCreateElement"},2591:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>s,default:()=>m,frontMatter:()=>i,metadata:()=>c,toc:()=>p});var o=n(3117),r=n(102),l=(n(7294),n(3905)),a=["components"],i={title:"\ud83d\udd79\ufe0f CLI"},s=void 0,c={unversionedId:"ecosystem/cli",id:"ecosystem/cli",title:"\ud83d\udd79\ufe0f CLI",description:"Selflow Cli depends on Selflow-Daemon if it doesn't run on your system, it will create it for you using Docker.",source:"@site/../../docs/ecosystem/cli.mdx",sourceDirName:"ecosystem",slug:"/ecosystem/cli",permalink:"/selflow/docs/ecosystem/cli",draft:!1,editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/../../docs/ecosystem/cli.mdx",tags:[],version:"current",frontMatter:{title:"\ud83d\udd79\ufe0f CLI"},sidebar:"tutorialSidebar",previous:{title:"Ecosystem",permalink:"/selflow/docs/ecosystem/"},next:{title:"\ud83d\udc31 Selflow-Daemon",permalink:"/selflow/docs/ecosystem/selflow-daemon"}},u={},p=[{value:"Installation",id:"installation",level:2},{value:"Commands",id:"commands",level:2},{value:"<code>recreate-daemon</code>",id:"recreate-daemon",level:3},{value:"Usage",id:"usage",level:4},{value:"<code>run</code>",id:"run",level:3},{value:"Usage",id:"usage-1",level:4},{value:"<code>status</code>",id:"status",level:3}],d={toc:p};function m(e){var t=e.components,n=(0,r.Z)(e,a);return(0,l.kt)("wrapper",(0,o.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("admonition",{type:"info"},(0,l.kt)("p",{parentName:"admonition"},"Selflow Cli depends on ",(0,l.kt)("a",{parentName:"p",href:"./selflow-daemon"},"Selflow-Daemon")," if it doesn't run on your system, it will create it for you using Docker.\nA Docker environment is also needed.")),(0,l.kt)("h2",{id:"installation"},"Installation"),(0,l.kt)("p",null,"Since Selflow is not ready for production yet, you will need to build the CLI yourself.\n",(0,l.kt)("strong",{parentName:"p"},"Don't worry"),", it only need to have ",(0,l.kt)("a",{parentName:"p",href:"https://go.dev/doc/install"},"Go")," setup on your machine in version ",(0,l.kt)("inlineCode",{parentName:"p"},"1.20")," or up and run this command :"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"go install github.com/selflow/selflow/apps/selflow-cli\n")),(0,l.kt)("p",null,"The executable will be added into your ",(0,l.kt)("inlineCode",{parentName:"p"},"$GOBIN")," location. Make sure to add it in your ",(0,l.kt)("inlineCode",{parentName:"p"},"$PATH"),"."),(0,l.kt)("p",null,"Check that the executable works with"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre"},"selflow-cli help\n")),(0,l.kt)("p",null,"If the CLI can not reach the Selflow-Daemon, it will create it using Docker."),(0,l.kt)("h2",{id:"commands"},"Commands"),(0,l.kt)("h3",{id:"recreate-daemon"},(0,l.kt)("inlineCode",{parentName:"h3"},"recreate-daemon")),(0,l.kt)("h4",{id:"usage"},"Usage"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"selflow-cli recreate-daemon\n")),(0,l.kt)("p",null,"Kill and start a new ",(0,l.kt)("a",{parentName:"p",href:"./selflow-daemon"},"Selflow Daemon")),(0,l.kt)("h3",{id:"run"},(0,l.kt)("inlineCode",{parentName:"h3"},"run")),(0,l.kt)("h4",{id:"usage-1"},"Usage"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"selflow-cli run ./path/to/your/workflow/file.json\n")),(0,l.kt)("p",null,"Start a run from the workflow of the given file. the file can be in json or yaml format. See Configuration reference for more details."),(0,l.kt)("p",null,"It currently shows logs in a debug format"),(0,l.kt)("h3",{id:"status"},(0,l.kt)("inlineCode",{parentName:"h3"},"status")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"selflow-cli status my-run-id\n")),(0,l.kt)("p",null,"Displays the status of each step of a run. The run id can be read from the ",(0,l.kt)("inlineCode",{parentName:"p"},"run")," command logs."))}m.isMDXComponent=!0}}]);