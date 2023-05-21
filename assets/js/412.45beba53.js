"use strict";(self.webpackChunkselflow=self.webpackChunkselflow||[]).push([[412],{9058:(e,t,a)=>{a.d(t,{Z:()=>Z});var r=a(102),l=a(7294),n=a(6010),i=a(215),o=a(7524),s=a(9960),m=a(5999);const c="sidebar_re4s",u="sidebarItemTitle_pO2u",d="sidebarItemList_Yudw",g="sidebarItem__DBe",v="sidebarItemLink_mo7H",p="sidebarItemLinkActive_I1ZP";function h(e){var t=e.sidebar;return l.createElement("aside",{className:"col col--3"},l.createElement("nav",{className:(0,n.Z)(c,"thin-scrollbar"),"aria-label":(0,m.I)({id:"theme.blog.sidebar.navAriaLabel",message:"Blog recent posts navigation",description:"The ARIA label for recent posts in the blog sidebar"})},l.createElement("div",{className:(0,n.Z)(u,"margin-bottom--md")},t.title),l.createElement("ul",{className:(0,n.Z)(d,"clean-list")},t.items.map((function(e){return l.createElement("li",{key:e.permalink,className:g},l.createElement(s.Z,{isNavLink:!0,to:e.permalink,className:v,activeClassName:p},e.title))})))))}var E=a(3102);function f(e){var t=e.sidebar;return l.createElement("ul",{className:"menu__list"},t.items.map((function(e){return l.createElement("li",{key:e.permalink,className:"menu__list-item"},l.createElement(s.Z,{isNavLink:!0,to:e.permalink,className:"menu__link",activeClassName:"menu__link--active"},e.title))})))}function b(e){return l.createElement(E.Zo,{component:f,props:e})}function N(e){var t=e.sidebar,a=(0,o.i)();return null!=t&&t.items.length?"mobile"===a?l.createElement(b,{sidebar:t}):l.createElement(h,{sidebar:t}):null}var _=["sidebar","toc","children"];function Z(e){var t=e.sidebar,a=e.toc,o=e.children,s=(0,r.Z)(e,_),m=t&&t.items.length>0;return l.createElement(i.Z,s,l.createElement("div",{className:"container margin-vert--lg"},l.createElement("div",{className:"row"},l.createElement(N,{sidebar:t}),l.createElement("main",{className:(0,n.Z)("col",{"col--7":m,"col--9 col--offset-1":!m}),itemScope:!0,itemType:"http://schema.org/Blog"},o),a&&l.createElement("div",{className:"col col--2"},a))))}},756:(e,t,a)=>{a.d(t,{Z:()=>H});var r=a(7294),l=a(6010),n=a(9460),i=a(4996);function o(e){var t,a=e.children,l=e.className,o=(0,n.C)(),s=o.frontMatter,m=o.assets,c=(0,i.C)().withBaseUrl,u=null!=(t=m.image)?t:s.image;return r.createElement("article",{className:l,itemProp:"blogPost",itemScope:!0,itemType:"http://schema.org/BlogPosting"},u&&r.createElement("meta",{itemProp:"image",content:c(u,{absolute:!0})}),a)}var s=a(9960);const m="title_f1Hy";function c(e){var t=e.className,a=(0,n.C)(),i=a.metadata,o=a.isBlogPostPage,c=i.permalink,u=i.title,d=o?"h1":"h2";return r.createElement(d,{className:(0,l.Z)(m,t),itemProp:"headline"},o?u:r.createElement(s.Z,{itemProp:"url",to:c},u))}var u=a(5999),d=a(2263),g=["zero","one","two","few","many","other"];function v(e){return g.filter((function(t){return e.includes(t)}))}var p={locale:"en",pluralForms:v(["one","other"]),select:function(e){return 1===e?"one":"other"}};function h(){var e=(0,d.Z)().i18n.currentLocale;return(0,r.useMemo)((function(){try{return t=e,a=new Intl.PluralRules(t),{locale:t,pluralForms:v(a.resolvedOptions().pluralCategories),select:function(e){return a.select(e)}}}catch(r){return console.error('Failed to use Intl.PluralRules for locale "'+e+'".\nDocusaurus will fallback to the default (English) implementation.\nError: '+r.message+"\n"),p}var t,a}),[e])}function E(){var e=h();return{selectMessage:function(t,a){return function(e,t,a){var r=e.split("|");if(1===r.length)return r[0];r.length>a.pluralForms.length&&console.error("For locale="+a.locale+", a maximum of "+a.pluralForms.length+" plural forms are expected ("+a.pluralForms.join(",")+"), but the message contains "+r.length+": "+e);var l=a.select(t),n=a.pluralForms.indexOf(l);return r[Math.min(n,r.length-1)]}(a,t,e)}}}const f="container_mt6G";function b(e){var t,a=e.readingTime,l=(t=E().selectMessage,function(e){var a=Math.ceil(e);return t(a,(0,u.I)({id:"theme.blog.post.readingTime.plurals",description:'Pluralized label for "{readingTime} min read". Use as much plural forms (separated by "|") as your language support (see https://www.unicode.org/cldr/cldr-aux/charts/34/supplemental/language_plural_rules.html)',message:"One min read|{readingTime} min read"},{readingTime:a}))});return r.createElement(r.Fragment,null,l(a))}function N(e){var t=e.date,a=e.formattedDate;return r.createElement("time",{dateTime:t,itemProp:"datePublished"},a)}function _(){return r.createElement(r.Fragment,null," \xb7 ")}function Z(e){var t=e.className,a=(0,n.C)().metadata,i=a.date,o=a.formattedDate,s=a.readingTime;return r.createElement("div",{className:(0,l.Z)(f,"margin-vert--md",t)},r.createElement(N,{date:i,formattedDate:o}),void 0!==s&&r.createElement(r.Fragment,null,r.createElement(_,null),r.createElement(b,{readingTime:s})))}function P(e){return e.href?r.createElement(s.Z,e):r.createElement(r.Fragment,null,e.children)}function k(e){var t=e.author,a=e.className,n=t.name,i=t.title,o=t.url,s=t.imageURL,m=t.email,c=o||m&&"mailto:"+m||void 0;return r.createElement("div",{className:(0,l.Z)("avatar margin-bottom--sm",a)},s&&r.createElement(P,{href:c,className:"avatar__photo-link"},r.createElement("img",{className:"avatar__photo",src:s,alt:n})),n&&r.createElement("div",{className:"avatar__intro",itemProp:"author",itemScope:!0,itemType:"https://schema.org/Person"},r.createElement("div",{className:"avatar__name"},r.createElement(P,{href:c,itemProp:"url"},r.createElement("span",{itemProp:"name"},n))),i&&r.createElement("small",{className:"avatar__subtitle",itemProp:"description"},i)))}const T="authorCol_Hf19",w="imageOnlyAuthorRow_pa_O",C="imageOnlyAuthorCol_G86a";function y(e){var t=e.className,a=(0,n.C)(),i=a.metadata.authors,o=a.assets;if(0===i.length)return null;var s=i.every((function(e){return!e.name}));return r.createElement("div",{className:(0,l.Z)("margin-top--md margin-bottom--sm",s?w:"row",t)},i.map((function(e,t){var a;return r.createElement("div",{className:(0,l.Z)(!s&&"col col--6",s?C:T),key:t},r.createElement(k,{author:Object.assign({},e,{imageURL:null!=(a=o.authorsImageUrls[t])?a:e.imageURL})}))})))}function B(){return r.createElement("header",null,r.createElement(c,null),r.createElement(Z,null),r.createElement(y,null))}var F=a(8780),I=a(2485);function x(e){var t=e.children,a=e.className,i=(0,n.C)().isBlogPostPage;return r.createElement("div",{id:i?F.blogPostContainerID:void 0,className:(0,l.Z)("markdown",a),itemProp:"articleBody"},r.createElement(I.Z,null,t))}var L=a(4881),M=a(6233),R=a(3117),U=a(102),A=["blogPostTitle"];function O(){return r.createElement("b",null,r.createElement(u.Z,{id:"theme.blog.post.readMore",description:"The label used in blog post item excerpts to link to full blog posts"},"Read More"))}function D(e){var t=e.blogPostTitle,a=(0,U.Z)(e,A);return r.createElement(s.Z,(0,R.Z)({"aria-label":(0,u.I)({message:"Read more about {title}",id:"theme.blog.post.readMoreLabel",description:"The ARIA label for the link to full blog posts from excerpts"},{title:t})},a),r.createElement(O,null))}const z="blogPostFooterDetailsFull_mRVl";function j(){var e=(0,n.C)(),t=e.metadata,a=e.isBlogPostPage,i=t.tags,o=t.title,s=t.editUrl,m=t.hasTruncateMarker,c=!a&&m,u=i.length>0;return u||c||s?r.createElement("footer",{className:(0,l.Z)("row docusaurus-mt-lg",a&&z)},u&&r.createElement("div",{className:(0,l.Z)("col",{"col--9":c})},r.createElement(M.Z,{tags:i})),a&&s&&r.createElement("div",{className:"col margin-top--sm"},r.createElement(L.Z,{editUrl:s})),c&&r.createElement("div",{className:(0,l.Z)("col text--right",{"col--3":u})},r.createElement(D,{blogPostTitle:o,to:t.permalink}))):null}function H(e){var t=e.children,a=e.className,i=(0,n.C)().isBlogPostPage?void 0:"margin-bottom--xl";return r.createElement(o,{className:(0,l.Z)(i,a)},r.createElement(B,null),r.createElement(x,null,t),r.createElement(j,null))}},4881:(e,t,a)=>{a.d(t,{Z:()=>d});var r=a(7294),l=a(5999),n=a(5281),i=a(3117),o=a(102),s=a(6010);const m="iconEdit_Z9Sw";var c=["className"];function u(e){var t=e.className,a=(0,o.Z)(e,c);return r.createElement("svg",(0,i.Z)({fill:"currentColor",height:"20",width:"20",viewBox:"0 0 40 40",className:(0,s.Z)(m,t),"aria-hidden":"true"},a),r.createElement("g",null,r.createElement("path",{d:"m34.5 11.7l-3 3.1-6.3-6.3 3.1-3q0.5-0.5 1.2-0.5t1.1 0.5l3.9 3.9q0.5 0.4 0.5 1.1t-0.5 1.2z m-29.5 17.1l18.4-18.5 6.3 6.3-18.4 18.4h-6.3v-6.2z"})))}function d(e){var t=e.editUrl;return r.createElement("a",{href:t,target:"_blank",rel:"noreferrer noopener",className:n.k.common.editThisPage},r.createElement(u,null),r.createElement(l.Z,{id:"theme.common.editThisPage",description:"The link label to edit the current page"},"Edit this page"))}},2244:(e,t,a)=>{a.d(t,{Z:()=>i});var r=a(7294),l=a(6010),n=a(9960);function i(e){var t=e.permalink,a=e.title,i=e.subLabel,o=e.isNext;return r.createElement(n.Z,{className:(0,l.Z)("pagination-nav__link",o?"pagination-nav__link--next":"pagination-nav__link--prev"),to:t},i&&r.createElement("div",{className:"pagination-nav__sublabel"},i),r.createElement("div",{className:"pagination-nav__label"},a))}},6233:(e,t,a)=>{a.d(t,{Z:()=>g});var r=a(7294),l=a(6010),n=a(5999),i=a(9960);const o="tag_zVej",s="tagRegular_sFm0",m="tagWithCount_h2kH";function c(e){var t=e.permalink,a=e.label,n=e.count;return r.createElement(i.Z,{href:t,className:(0,l.Z)(o,n?m:s)},a,n&&r.createElement("span",null,n))}const u="tags_jXut",d="tag_QGVx";function g(e){var t=e.tags;return r.createElement(r.Fragment,null,r.createElement("b",null,r.createElement(n.Z,{id:"theme.tags.tagsListLabel",description:"The label alongside a tag list"},"Tags:")),r.createElement("ul",{className:(0,l.Z)(u,"padding--none","margin-left--sm")},t.map((function(e){var t=e.label,a=e.permalink;return r.createElement("li",{key:a,className:d},r.createElement(c,{label:t,permalink:a}))}))))}},9460:(e,t,a)=>{a.d(t,{C:()=>o,n:()=>i});var r=a(7294),l=a(4700),n=r.createContext(null);function i(e){var t=e.children,a=e.content,l=e.isBlogPostPage,i=function(e){var t=e.content,a=e.isBlogPostPage;return(0,r.useMemo)((function(){return{metadata:t.metadata,frontMatter:t.frontMatter,assets:t.assets,toc:t.toc,isBlogPostPage:a}}),[t,a])}({content:a,isBlogPostPage:void 0!==l&&l});return r.createElement(n.Provider,{value:i},t)}function o(){var e=(0,r.useContext)(n);if(null===e)throw new l.i6("BlogPostProvider");return e}}}]);