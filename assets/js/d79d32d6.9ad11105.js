"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[154],{3905:(e,t,a)=>{a.d(t,{Zo:()=>p,kt:()=>h});var n=a(7294);function l(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function o(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function r(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?o(Object(a),!0).forEach((function(t){l(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):o(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function i(e,t){if(null==e)return{};var a,n,l=function(e,t){if(null==e)return{};var a,n,l={},o=Object.keys(e);for(n=0;n<o.length;n++)a=o[n],t.indexOf(a)>=0||(l[a]=e[a]);return l}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)a=o[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(l[a]=e[a])}return l}var s=n.createContext({}),c=function(e){var t=n.useContext(s),a=t;return e&&(a="function"==typeof e?e(t):r(r({},t),e)),a},p=function(e){var t=c(e.components);return n.createElement(s.Provider,{value:t},e.children)},u="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var a=e.components,l=e.mdxType,o=e.originalType,s=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),u=c(a),f=l,h=u["".concat(s,".").concat(f)]||u[f]||m[f]||o;return a?n.createElement(h,r(r({ref:t},p),{},{components:a})):n.createElement(h,r({ref:t},p))}));function h(e,t){var a=arguments,l=t&&t.mdxType;if("string"==typeof e||l){var o=a.length,r=new Array(o);r[0]=f;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[u]="string"==typeof e?e:l,r[1]=i;for(var c=2;c<o;c++)r[c]=a[c];return n.createElement.apply(null,r)}return n.createElement.apply(null,a)}f.displayName="MDXCreateElement"},7942:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>p,contentTitle:()=>s,default:()=>f,frontMatter:()=>i,metadata:()=>c,toc:()=>u});var n=a(3117),l=a(102),o=(a(7294),a(3905)),r=["components"],i={slug:"state-of-selflow",title:"The State of Selflow",authors:"anthony-quere",tags:[]},s=void 0,c={permalink:"/selflow/blog/state-of-selflow",editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/blog/2022-01-31-state-of-selflow.md",source:"@site/blog/2022-01-31-state-of-selflow.md",title:"The State of Selflow",description:"Selflow was created as a school project should have a consistent development over the next months.",date:"2022-01-31T00:00:00.000Z",formattedDate:"January 31, 2022",tags:[],readingTime:.72,hasTruncateMarker:!1,authors:[{name:"Anthony Qu\xe9r\xe9",title:"Creator of Selflow",url:"https://github.com/Anthony-Jhoiro",imageURL:"https://github.com/Anthony-Jhoiro.png",key:"anthony-quere"}],frontMatter:{slug:"state-of-selflow",title:"The State of Selflow",authors:"anthony-quere",tags:[]},nextItem:{title:"What is Selflow ?",permalink:"/selflow/blog/what-is-selflow"}},p={authorsImageUrls:[void 0]},u=[{value:"Objectives for May 14th",id:"objectives-for-may-14th",level:2},{value:"If I got the time",id:"if-i-got-the-time",level:2}],m={toc:u};function f(e){var t=e.components,a=(0,l.Z)(e,r);return(0,o.kt)("wrapper",(0,n.Z)({},m,a,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("p",null,"Selflow was created as a school project should have a consistent development over the next months."),(0,o.kt)("h2",{id:"objectives-for-may-14th"},"Objectives for May 14th"),(0,o.kt)("ul",{className:"contains-task-list"},(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Run a workflow from the ",(0,o.kt)("inlineCode",{parentName:"li"},"selflow-core")," to the ",(0,o.kt)("inlineCode",{parentName:"li"},"selflow-runner")," with plugins"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Be able to parallelize steps"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Build custom plugins for custom workflow steps"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Build custom plugins to trigger workflows"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Be able to run commands inside containers with the specified image"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Execute commands on a remote machine with SSH"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Strong integration with GitHub for CI/CD workflows (if possible in self-hosted)"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Visualize steps using a UI")),(0,o.kt)("h2",{id:"if-i-got-the-time"},"If I got the time"),(0,o.kt)("ul",{className:"contains-task-list"},(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Build a Sass platform to manage multiple run configurations (might be required for the full GitHub integration)"),(0,o.kt)("li",{parentName:"ul",className:"task-list-item"},(0,o.kt)("input",{parentName:"li",type:"checkbox",checked:!1,disabled:!0})," ","Run workflows on public selflow instances")))}f.isMDXComponent=!0}}]);