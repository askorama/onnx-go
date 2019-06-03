var msg = new SpeechSynthesisUtterance();
var voices = window.speechSynthesis.getVoices();
//msg.voice = voices[3]; // Note: some voices don't support altering params
//msg.voiceURI = 'native';
//msg.volume = 1; // 0 to 1
//msg.rate = 0.7; // 0.1 to 10
//msg.pitch = 1.5; //0 to 2
//msg.lang = 'en-US';

msg.onend = function(e) {
  console.log('Finished in ' + event.elapsedTime + ' seconds.');
};


var canvasMnist;
var image;
var bThick = 20;
var cColour = '#FFFFFF'; // #black

//Initialize  fabric
function initcanvasMnist(){
  canvasMnist = new fabric.Canvas('canvasMnist');
  canvasMnist.isDrawingMode = true;
  canvasMnist.freeDrawingBrush.width = bThick;
  canvasMnist.freeDrawingBrush.color = cColour;
  canvasMnist.backgroundColor = '#000000';
}


//reset canvasMnist
function reset() {
  canvasMnist.clear();
  document.getElementById("infoMnist").innerHTML = "...";
}

function postImage() {
  const dataURL = canvasMnist.toDataURL();
  $.ajax({
    type: 'POST',
    url: '/image',
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    data: JSON.stringify({image: dataURL, model: "mnist"}),
    success: function(data){
      console.log(data);
      $( "#infoMnist" ).html( data[0].Result );
      msg.text = data[0].Result;
      speechSynthesis.speak(msg);
    },
    failure: function(errMsg) {
      alert(errMsg);
    }
  })
}
async function uploadModelMnist(e)
{
  let name = { modelname:'mnist' };
  let formData = new FormData();
  let model = e.files[0];

  formData.append("model", model);
  formData.append("name", JSON.stringify(name));

  try {
    //let r = await fetch('/model', {method: "POST", body: formData});
    let r = await fetch('/model', {method: "POST", body: e.files[0]});
    document.getElementById("gomnist").disabled = false;

    console.log('HTTP response code:',r.status);
  } catch(e) {
    console.log('Huston we have problem...:', e);
  }
  initcanvasMnist();
}

