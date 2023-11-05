"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[346],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>k});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=r.createContext({}),p=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},c=function(e){var t=p(e.components);return r.createElement(l.Provider,{value:t},e.children)},d="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,l=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),d=p(n),u=o,k=d["".concat(l,".").concat(u)]||d[u]||m[u]||a;return n?r.createElement(k,i(i({ref:t},c),{},{components:n})):r.createElement(k,i({ref:t},c))}));function k(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,i=new Array(a);i[0]=u;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s[d]="string"==typeof e?e:o,i[1]=s;for(var p=2;p<a;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},9949:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>u,frontMatter:()=>s,metadata:()=>p,toc:()=>d});var r=n(3117),o=n(102),a=(n(7294),n(3905)),i=["components"],s={title:"\ud83d\udc0b Docker Steps"},l=void 0,p={unversionedId:"steps/docker",id:"steps/docker",title:"\ud83d\udc0b Docker Steps",description:"Docker steps are used to execute code inside a docker container.",source:"@site/../../docs/steps/docker.md",sourceDirName:"steps",slug:"/steps/docker",permalink:"/selflow/docs/steps/docker",draft:!1,editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/../../docs/steps/docker.md",tags:[],version:"current",frontMatter:{title:"\ud83d\udc0b Docker Steps"},sidebar:"tutorialSidebar",previous:{title:"Steps",permalink:"/selflow/docs/steps/"},next:{title:"\u270f\ufe0f Workflow Syntax",permalink:"/selflow/docs/workflow-syntax"}},c={},d=[{value:"Options",id:"options",level:2},{value:"<code>image</code> (<strong>Required</strong>)",id:"image",level:3},{value:"<code>commands</code> (<strong>Required</strong>)",id:"commands",level:3},{value:"<code>persistence</code>",id:"persistence",level:3}],m={toc:d};function u(e){var t=e.components,n=(0,o.Z)(e,i);return(0,a.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("p",null,"Docker steps are used to execute code inside a docker container."),(0,a.kt)("h2",{id:"options"},"Options"),(0,a.kt)("h3",{id:"image"},(0,a.kt)("inlineCode",{parentName:"h3"},"image")," (",(0,a.kt)("strong",{parentName:"h3"},"Required"),")"),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},"Supports Go Template")),(0,a.kt)("p",null,"Docker image to use as a base."),(0,a.kt)("admonition",{type:"warning"},(0,a.kt)("p",{parentName:"admonition"},"There is currently a limitation and the user of the image must be a root user, otherwise the commands won't be executed.\nthe image also need to have a shell at ",(0,a.kt)("inlineCode",{parentName:"p"},"/bin/sh"))),(0,a.kt)("h3",{id:"commands"},(0,a.kt)("inlineCode",{parentName:"h3"},"commands")," (",(0,a.kt)("strong",{parentName:"h3"},"Required"),")"),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},"Supports Go Template")),(0,a.kt)("p",null,"Commands to run on the container. For now, that will always be executed by a shell located at ",(0,a.kt)("inlineCode",{parentName:"p"},"/bin/sh"),".\nyou can use a multiline string fot that field"),(0,a.kt)("h3",{id:"persistence"},(0,a.kt)("inlineCode",{parentName:"h3"},"persistence")),(0,a.kt)("p",null,"The persistence attribute allows step to share files between them using docker volumes.\nThis way, if a file is added or change in a step, it can also be available in another step ",(0,a.kt)("strong",{parentName:"p"},"as long as it is a Docker Step"),"."),(0,a.kt)("p",null,"The volumes are scoped to a specific run and can't be used in other runs."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},'workflow:\n  timeout: 5m\n  steps:\n    step-a:\n      kind: docker\n      with:\n        persistence:\n          my-volume: /workdir\n        image: node:lts\n        commands: |\n          echo "Hello!" > /workdir/someFile.txt\n\n    step-b:\n      kind: docker\n      needs:\n        - step-a\n      with:\n        persistence:\n          my-volume: /workdir\n        image: node:lts\n        commands: |\n          cat /workdir/someFile.txt\n')),(0,a.kt)("p",null,"In this example, ",(0,a.kt)("inlineCode",{parentName:"p"},"step-a")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"step-b"),"depends on the same persistence volume called ",(0,a.kt)("inlineCode",{parentName:"p"},"my-volume"),". In both case, it will\nbe mapped to the ",(0,a.kt)("inlineCode",{parentName:"p"},"/workdir")," path on the container so if a step add a file in this directory, it will be available for the other step.\nUsing the ",(0,a.kt)("a",{parentName:"p",href:"./#needs"},(0,a.kt)("inlineCode",{parentName:"a"},"needs"))," attribute, we specify that ",(0,a.kt)("inlineCode",{parentName:"p"},"step-a")," will be executed before ",(0,a.kt)("inlineCode",{parentName:"p"},"step-b"),".\nAfter ",(0,a.kt)("inlineCode",{parentName:"p"},"step-a")," wrote in ",(0,a.kt)("inlineCode",{parentName:"p"},"/workdir/someFile.txt")," and finished, the file will be available to be read by ",(0,a.kt)("inlineCode",{parentName:"p"},"step-b"),"."),(0,a.kt)("p",null,"Here, ",(0,a.kt)("inlineCode",{parentName:"p"},"step-b")," will log ",(0,a.kt)("inlineCode",{parentName:"p"},"Hello!")))}u.isMDXComponent=!0}}]);