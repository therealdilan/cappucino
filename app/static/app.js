var colorPicker, resetCode;

window.onload = function() {
  colorPicker = document.getElementById("colorPicker");
  resetCode = document.getElementById("resetCode");

  colorPicker.addEventListener("change", watchColorPicker, false);
  resetCode.addEventListener("click", clearCode, false);
}

function watchColorPicker(event) {
  document.getElementById("preview").style.backgroundColor = event.target.value;
}

function clearCode() {
  document.getElementById("code").value = "";
  document.getElementById("render").innerHTML = "";
}
