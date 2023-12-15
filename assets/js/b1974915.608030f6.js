"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[868],{871:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>s,default:()=>f,frontMatter:()=>i,metadata:()=>c,toc:()=>u});var r=n(3117),o=n(102),a=(n(7294),n(3905)),l=["components"],i={title:"Contributing"},s=void 0,c={unversionedId:"contributing",id:"contributing",title:"Contributing",description:"Pull Requests are welcome !",source:"@site/../../docs/contributing.md",sourceDirName:".",slug:"/contributing",permalink:"/selflow/docs/contributing",draft:!1,editUrl:"https://github.com/selflow/selflow/edit/main/apps/selflow-documentation/../../docs/contributing.md",tags:[],version:"current",frontMatter:{title:"Contributing"},sidebar:"tutorialSidebar",previous:{title:"Introduction",permalink:"/selflow/docs/intro"},next:{title:"Ecosystem",permalink:"/selflow/docs/ecosystem/"}},p={},u=[{value:"Set up the project locally",id:"set-up-the-project-locally",level:2},{value:"Requirements",id:"requirements",level:3},{value:"Repository Structure",id:"repository-structure",level:2}],m={toc:u};function f(e){var t=e.components,n=(0,o.Z)(e,l);return(0,a.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("p",null,"Pull Requests are welcome !\nIf you have an idea or want to start a major feature, please ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/selflow/selflow/discussions/new/choose"},"open a discussion")," on GitHub."),(0,a.kt)("h2",{id:"set-up-the-project-locally"},"Set up the project locally"),(0,a.kt)("h3",{id:"requirements"},"Requirements"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Node.js"),(0,a.kt)("li",{parentName:"ul"},"Yarn"),(0,a.kt)("li",{parentName:"ul"},"Go")),(0,a.kt)("p",null,"Depending on what you are doing, you might want to have Docker installed and set up locally."),(0,a.kt)("p",null,"Setting up the project locally is fairly easy. Start by cloning the repository :"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"git clone https://github.com/selflow/selflow.git\ncd selflow\n")),(0,a.kt)("p",null,"Now you can install NX dependencies using ",(0,a.kt)("inlineCode",{parentName:"p"},"yarn")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"yarn\n")),(0,a.kt)("p",null,"If needed, you can also install go dependencies using"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"go mod download\n")),(0,a.kt)("admonition",{type:"note"},(0,a.kt)("p",{parentName:"admonition"},"If you want to change something to the selflow core, you do not have to set up NX on your repository but it will make your life easier")),(0,a.kt)("h2",{id:"repository-structure"},"Repository Structure"),(0,a.kt)("p",null,"The Selflow repository is a mono-repository that is meant to manage the majority of the Selflow ecosystem.\nIt allows us to have a global vision of the impact pof any new feature."),(0,a.kt)("p",null,"For that we are using ",(0,a.kt)("a",{parentName:"p",href:"https://nx.dev/"},"NX")," which is an amazing tool !"),(0,a.kt)("p",null,"The main directories are :"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},".\n\u251c\u2500\u2500 apps              # Applications and e2e tests for those applications\n\u251c\u2500\u2500 libs\n\u2502   \u251c\u2500\u2500 core          # Selflow core-libraries\n\u2502   \u251c\u2500\u2500 [app-name]    # Libraries specifics to [app-name]\n\u2502   \u251c\u2500\u2500 protos        # Protobufs\n\u2502   \u2514\u2500\u2500 ui            # Libraries shared by all UI projects\n\u251c\u2500\u2500 docs              # Markdown documentation\n\u2514\u2500\u2500 assets            # Assets for the main README\n")))}f.isMDXComponent=!0},3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>d});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},u="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},f=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,s=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),u=c(n),f=o,d=u["".concat(s,".").concat(f)]||u[f]||m[f]||a;return n?r.createElement(d,l(l({ref:t},p),{},{components:n})):r.createElement(d,l({ref:t},p))}));function d(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,l=new Array(a);l[0]=f;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[u]="string"==typeof e?e:o,l[1]=i;for(var c=2;c<a;c++)l[c]=n[c];return r.createElement.apply(null,l)}return r.createElement.apply(null,n)}f.displayName="MDXCreateElement"}}]);