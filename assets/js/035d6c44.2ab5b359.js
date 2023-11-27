"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[307],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var i=r.createContext({}),c=function(e){var t=r.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(i.Provider,{value:t},e.children)},u="mdxType",f={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,l=e.originalType,i=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),u=c(n),d=o,m=u["".concat(i,".").concat(d)]||u[d]||f[d]||l;return n?r.createElement(m,a(a({ref:t},p),{},{components:n})):r.createElement(m,a({ref:t},p))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var l=n.length,a=new Array(l);a[0]=d;var s={};for(var i in t)hasOwnProperty.call(t,i)&&(s[i]=t[i]);s.originalType=e,s[u]="string"==typeof e?e:o,a[1]=s;for(var c=2;c<l;c++)a[c]=n[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},8266:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>i,default:()=>d,frontMatter:()=>s,metadata:()=>c,toc:()=>u});var r=n(3117),o=n(102),l=(n(7294),n(3905)),a=["components"],s={slug:"run",title:"\u2328 run"},i="\u2328\ufe0f `run`",c={unversionedId:"ecosystem/cli/selflow_run",id:"ecosystem/cli/selflow_run",title:"\u2328 run",description:"selflow run",source:"@site/../../docs/ecosystem/cli/selflow_run.md",sourceDirName:"ecosystem/cli",slug:"/ecosystem/cli/run",permalink:"/selflow/docs/ecosystem/cli/run",draft:!1,editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/../../docs/ecosystem/cli/selflow_run.md",tags:[],version:"current",frontMatter:{slug:"run",title:"\u2328 run"},sidebar:"tutorialSidebar",previous:{title:"\u2328 recreate-daemon",permalink:"/selflow/docs/ecosystem/cli/recreate-daemon"},next:{title:"\u2328 status",permalink:"/selflow/docs/ecosystem/cli/status"}},p={},u=[{value:"selflow run",id:"selflow-run",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3},{value:"Auto generated by spf13/cobra on 27-Nov-2023",id:"auto-generated-by-spf13cobra-on-27-nov-2023",level:6}],f={toc:u};function d(e){var t=e.components,n=(0,o.Z)(e,a);return(0,l.kt)("wrapper",(0,r.Z)({},f,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"\ufe0f-run"},"\u2328\ufe0f ",(0,l.kt)("inlineCode",{parentName:"h1"},"run")),(0,l.kt)("h2",{id:"selflow-run"},"selflow run"),(0,l.kt)("p",null,"Start a workflow on the Selflow-Daemon and wait for the end of its execution."),(0,l.kt)("h3",{id:"synopsis"},"Synopsis"),(0,l.kt)("p",null,"Start a workflow on the Selflow-Daemon and wait for the end of its execution."),(0,l.kt)("p",null,"If the command is stopped, the workflow will not be stopped. The commands also gives logs about what is happening."),(0,l.kt)("p",null,"The workflow file must follow the selflow workflow syntax (",(0,l.kt)("a",{parentName:"p",href:"https://selflow.github.io/selflow/docs/workflow-syntax"},"https://selflow.github.io/selflow/docs/workflow-syntax"),")."),(0,l.kt)("p",null,"It can be written using ",(0,l.kt)("strong",{parentName:"p"},"YAML")," or ",(0,l.kt)("strong",{parentName:"p"},"JSON"),"."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre"},"selflow run ./filename [flags]\n")),(0,l.kt)("h3",{id:"options"},"Options"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre"},"  -h, --help   help for run\n")),(0,l.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre"},"      --recreate-daemon   Kill and recreate the daemon if it already exists\n")),(0,l.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"/selflow/docs/ecosystem/cli/selflow"},"selflow")," - Selflow is a workflow orchestration tool")),(0,l.kt)("h6",{id:"auto-generated-by-spf13cobra-on-27-nov-2023"},"Auto generated by spf13/cobra on 27-Nov-2023"))}d.isMDXComponent=!0}}]);