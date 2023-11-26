"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[689],{3905:(e,t,o)=>{o.d(t,{Zo:()=>p,kt:()=>d});var n=o(7294);function l(e,t,o){return t in e?Object.defineProperty(e,t,{value:o,enumerable:!0,configurable:!0,writable:!0}):e[t]=o,e}function r(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),o.push.apply(o,n)}return o}function i(e){for(var t=1;t<arguments.length;t++){var o=null!=arguments[t]?arguments[t]:{};t%2?r(Object(o),!0).forEach((function(t){l(e,t,o[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):r(Object(o)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(o,t))}))}return e}function s(e,t){if(null==e)return{};var o,n,l=function(e,t){if(null==e)return{};var o,n,l={},r=Object.keys(e);for(n=0;n<r.length;n++)o=r[n],t.indexOf(o)>=0||(l[o]=e[o]);return l}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(n=0;n<r.length;n++)o=r[n],t.indexOf(o)>=0||Object.prototype.propertyIsEnumerable.call(e,o)&&(l[o]=e[o])}return l}var a=n.createContext({}),c=function(e){var t=n.useContext(a),o=t;return e&&(o="function"==typeof e?e(t):i(i({},t),e)),o},p=function(e){var t=c(e.components);return n.createElement(a.Provider,{value:t},e.children)},m="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var o=e.components,l=e.mdxType,r=e.originalType,a=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),m=c(o),f=l,d=m["".concat(a,".").concat(f)]||m[f]||u[f]||r;return o?n.createElement(d,i(i({ref:t},p),{},{components:o})):n.createElement(d,i({ref:t},p))}));function d(e,t){var o=arguments,l=t&&t.mdxType;if("string"==typeof e||l){var r=o.length,i=new Array(r);i[0]=f;var s={};for(var a in t)hasOwnProperty.call(t,a)&&(s[a]=t[a]);s.originalType=e,s[m]="string"==typeof e?e:l,i[1]=s;for(var c=2;c<r;c++)i[c]=o[c];return n.createElement.apply(null,i)}return n.createElement.apply(null,o)}f.displayName="MDXCreateElement"},9254:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>p,contentTitle:()=>a,default:()=>f,frontMatter:()=>s,metadata:()=>c,toc:()=>m});var n=o(3117),l=o(102),r=(o(7294),o(3905)),i=["components"],s={slug:"completion/bash",title:"\u2328 completion_bash"},a="\u2328\ufe0f `completion_bash`",c={unversionedId:"ecosystem/cli/selflow_completion_bash",id:"ecosystem/cli/selflow_completion_bash",title:"\u2328 completion_bash",description:"selflow completion bash",source:"@site/../../docs/ecosystem/cli/selflow_completion_bash.md",sourceDirName:"ecosystem/cli",slug:"/ecosystem/cli/completion/bash",permalink:"/selflow/docs/ecosystem/cli/completion/bash",draft:!1,editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/../../docs/ecosystem/cli/selflow_completion_bash.md",tags:[],version:"current",frontMatter:{slug:"completion/bash",title:"\u2328 completion_bash"},sidebar:"tutorialSidebar",previous:{title:"\u2328 completion",permalink:"/selflow/docs/ecosystem/cli/completion"},next:{title:"\u2328 completion_fish",permalink:"/selflow/docs/ecosystem/cli/completion/fish"}},p={},m=[{value:"selflow completion bash",id:"selflow-completion-bash",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Linux:",id:"linux",level:4},{value:"macOS:",id:"macos",level:4},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3},{value:"Auto generated by spf13/cobra on 26-Nov-2023",id:"auto-generated-by-spf13cobra-on-26-nov-2023",level:6}],u={toc:m};function f(e){var t=e.components,o=(0,l.Z)(e,i);return(0,r.kt)("wrapper",(0,n.Z)({},u,o,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"\ufe0f-completion_bash"},"\u2328\ufe0f ",(0,r.kt)("inlineCode",{parentName:"h1"},"completion_bash")),(0,r.kt)("h2",{id:"selflow-completion-bash"},"selflow completion bash"),(0,r.kt)("p",null,"Generate the autocompletion script for bash"),(0,r.kt)("h3",{id:"synopsis"},"Synopsis"),(0,r.kt)("p",null,"Generate the autocompletion script for the bash shell."),(0,r.kt)("p",null,"This script depends on the 'bash-completion' package.\nIf it is not installed already, you can install it via your OS's package manager."),(0,r.kt)("p",null,"To load completions in your current shell session:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"source <(selflow completion bash)\n")),(0,r.kt)("p",null,"To load completions for every new session, execute once:"),(0,r.kt)("h4",{id:"linux"},"Linux:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"selflow completion bash > /etc/bash_completion.d/selflow\n")),(0,r.kt)("h4",{id:"macos"},"macOS:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"selflow completion bash > $(brew --prefix)/etc/bash_completion.d/selflow\n")),(0,r.kt)("p",null,"You will need to start a new shell for this setup to take effect."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"selflow completion bash\n")),(0,r.kt)("h3",{id:"options"},"Options"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"  -h, --help              help for bash\n      --no-descriptions   disable completion descriptions\n")),(0,r.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"      --recreate-daemon   Kill and recreate the daemon if it already exists\n")),(0,r.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"/selflow/docs/ecosystem/cli/completion"},"selflow completion")," - Generate the autocompletion script for the specified shell")),(0,r.kt)("h6",{id:"auto-generated-by-spf13cobra-on-26-nov-2023"},"Auto generated by spf13/cobra on 26-Nov-2023"))}f.isMDXComponent=!0}}]);