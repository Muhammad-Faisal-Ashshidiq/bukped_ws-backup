import { setInner, addInner } from "https://jscroot.github.io/element/croot.js";
import {get} from "https://jscroot.github.io/api/croot.js"; 
import { rowtabel } from "./view/tabel.js";
setInner("namadivisi","Dari croot.js");
get("https://gocroot-baru.herokuapp.com/hd", tabelContent);
 
var isi = {};

function tabelContent(result){
    userTable(result);
    isi = result;
}

function userTable(jsonParse){

    jsonParse.forEach((element, index) => {
    let row = rowtabel.replace("#Nama#", element.sistemcomp).replace("#phonenumber#", element.status);
    console.log(element._id);
    addInner("demo",row);
    });
}