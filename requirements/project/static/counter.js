const textArea = document.getElementById('str-input');
const inputCounter = document.getElementById('inputCounter');

textArea.addEventListener('input',function() {
    const counter = textArea.nodeValue.length;
    inputCounter.textContent = `${counter} character${counter !==1 ? 's' : ''}`
});