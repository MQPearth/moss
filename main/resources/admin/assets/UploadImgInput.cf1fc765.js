import{u as y}from"./vue-request.6886b8d7.js";import{l as S}from"./index.3eae3a9f.js";import{t as w,r as c,w as m,j as l,k as x,l as B,m as s,v as i,q as C,n as I}from"./@vue.0987707a.js";const N={__name:"UploadImgInput",props:{modelValue:String,inputStyle:String},emits:["update:modelValue"],setup(p,{emit:d}){const f=p,{modelValue:o}=w(f),t=c(o.value),n=c();m(o,()=>{t.value=o.value}),m(t,()=>{d("update:modelValue",t.value)});const{run:_}=y(S,{manual:!0,onBefore:()=>{},onSuccess:a=>{a.success&&a.data.length>0&&(t.value=a.data[0])}});function v(a){let e=new FormData;for(let u of a.target.files)e.append("file",u);_(e),n.value.value=""}return(a,e)=>{const u=l("icon-upload"),g=l("a-link"),V=l("a-input"),k=l("a-input-group");return x(),B(k,null,{default:s(()=>[i(V,{style:C(p.inputStyle),modelValue:t.value,"onUpdate:modelValue":e[1]||(e[1]=r=>t.value=r),"allow-clear":""},{suffix:s(()=>[i(g,{onClick:e[0]||(e[0]=r=>n.value.click()),hoverable:!0},{default:s(()=>[i(u)]),_:1})]),_:1},8,["style","modelValue"]),I("input",{class:"hidden",type:"file",ref_key:"fileInput",ref:n,accept:"image/*",onChange:v},null,544)]),_:1})}}};export{N as _};
