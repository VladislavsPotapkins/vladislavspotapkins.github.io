let i = 0;
let arr = [0,0,0,0,0,0,0,0,0]

function result(){

      if (arr[0]===arr[1] && arr[0]===arr[2] && arr[0] !== 0) { if (arr[0]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[0]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[3]===arr[4] && arr[3]===arr[5] && arr[3] !== 0) { if (arr[3]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[3]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[6]===arr[7] && arr[6]===arr[8] && arr[6] !== 0) { if (arr[6]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[6]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[0]===arr[3] && arr[0]===arr[6] && arr[0] !== 0) { if (arr[0]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[0]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[1]===arr[4] && arr[1]===arr[7] && arr[1] !== 0) { if (arr[1]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[1]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[2]===arr[5] && arr[2]===arr[8] && arr[2] !== 0) { if (arr[2]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[2]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[0]===arr[4] && arr[0]===arr[8] && arr[0] !== 0) { if (arr[0]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[0]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (arr[2]===arr[4] && arr[2]===arr[6] && arr[2] !== 0) { if (arr[2]===1){alert ( "Player 1 Wins!");alert("Game Over!!!");} else if (arr[2]===2){alert ( "Player 2 Wins!");alert("Game Over!!!");}}
      if (i>7) {alert("Game Over!!!")}
}


document.getElementById('1').onclick = function() {
document.getElementById('1').hidden = true;
if (i%2===0){
document.getElementById('11').hidden = false;
arr[0] = 1;
result();
} else {
document.getElementById('12').hidden = false;
arr[0] = 2;
result();
}
i++;
}

document.getElementById('2').onclick = function() {
document.getElementById('2').hidden = true;
if (i%2===0){
document.getElementById('21').hidden = false;
arr[1] = 1;
result();
} else {
document.getElementById('22').hidden = false;
arr[1] = 2;
result();
}
i++;
}

document.getElementById('3').onclick = function() {
document.getElementById('3').hidden = true;
if (i%2===0){
document.getElementById('31').hidden = false;
arr[2] = 1;
result();
} else {
document.getElementById('32').hidden = false;
arr[2] = 2;
result();
}
i++;    }

document.getElementById('4').onclick = function() {
document.getElementById('4').hidden = true;
if (i%2===0){
document.getElementById('41').hidden = false;
arr[3] = 1;
result();
} else {
document.getElementById('42').hidden = false;
arr[3] = 2;
result();
}
i++;    }

document.getElementById('5').onclick = function() {
document.getElementById('5').hidden = true;
if (i%2===0){
document.getElementById('51').hidden = false;
arr[4] = 1;
result();
} else {
document.getElementById('52').hidden = false;
arr[4] = 2;
result();
}
i++;    }

document.getElementById('6').onclick = function() {
document.getElementById('6').hidden = true;
if (i%2===0){
document.getElementById('61').hidden = false;
arr[5] = 1;
result();
} else {
document.getElementById('62').hidden = false;
arr[5] = 2;
result();
}
i++;    }

document.getElementById('7').onclick = function() {
document.getElementById('7').hidden = true;
if (i%2===0){
document.getElementById('71').hidden = false;
arr[6] = 1;
result();
} else {
document.getElementById('72').hidden = false;
arr[6] = 2;
result();
}
i++;    }

document.getElementById('8').onclick = function() {
document.getElementById('8').hidden = true;
if (i%2===0){
document.getElementById('81').hidden = false;
arr[7] = 1;
result();
} else {
document.getElementById('82').hidden = false;
arr[7] = 2;
result();
}
i++;    }

document.getElementById('9').onclick = function() {
document.getElementById('9').hidden = true;
if (i%2===0){
document.getElementById('91').hidden = false;
arr[8] = 1;
result();
} else {
document.getElementById('92').hidden = false;
arr[8] = 2;
result();
}

i++;    }
