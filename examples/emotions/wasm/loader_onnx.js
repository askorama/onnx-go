//reset canvas
function reset() {           
   document.getElementById("info").innerHTML = "...";
}

if (!WebAssembly.instantiateStreaming) { // polyfill
         WebAssembly.instantiateStreaming = async (resp, importObject) => {
                 const source = await (await resp).arrayBuffer();
                 return await WebAssembly.instantiate(source, importObject);
       };
}

  const go = new Go();
  let mod, inst;
async function load() {
    document.getElementById("loadButton").disabled = true;
    document.getElementById("loadButton").style.color = "grey";
    document.getElementById("info").innerHTML = "Downloading WASM file...";
    const fetchPromise = fetch('emotions.wasm');
    document.getElementById("info").innerHTML = "Compiling WASM file, please wait...";
    WebAssembly.instantiateStreaming(fetchPromise, go.importObject).then((result) => {
         mod = result.module;
         inst = result.instance;
         document.getElementById("runButton").disabled = false;
    document.getElementById("info").innerHTML = "Ready to run...";
         document.getElementById("runButton").style.color = "green";
  });
}

async function run() {
    document.getElementById("runButton").disabled = true;
    console.clear();
    document.getElementById("info").innerHTML = "Starting...";
    await go.run(inst);
  //inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
    inst = await WebAssembly.instantiate(mod, go.importObject).catch(error => {
      document.getElementById("info").innerHTML = "program failed (out-of-memory?)<br>(this program does not run properly on mobile)<br>reload the page to retry";
         document.getElementById("info").style.color = "red";
    });
}
function loadFile(){
    var x = document.getElementById("knowledgeFile");
    var txt = "";
    if ('files' in x) {
        if (x.files.length == 0) {
            txt = "Select one or more files.";
        } else {
            for (var i = 0; i < x.files.length; i++) {
                txt += "<br><strong>" + (i+1) + ". file</strong><br>";
                var file = x.files[i];
                if ('name' in file) {
                    txt += "name: " + file.name + "<br>";
                }
                if ('size' in file) {
                    txt += "size: " + file.size + " bytes <br>";
                }
            }
        }
    }
    else {
        if (x.value == "") {
            txt += "Select one or more files.";
        } else {
            txt += "The files property is not supported by your browser!";
            txt  += "<br>The path of the selected file: " + x.value; // If the browser does not support the files property, it will return the path of the selected file instead.
        }
    }
    document.getElementById("demo").innerHTML = txt;
   document.getElementById("loadButton").disabled = false;
         document.getElementById("loadButton").style.color = "green";
}

function fileOK() {
   document.getElementById("loadButton").disabled = false;
   document.getElementById("loadButton").style.color = "green";
}
