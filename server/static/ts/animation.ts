const canvas = document.getElementById("title") as HTMLCanvasElement;
const howMany = 60
const randomPhrases = ["The CIA is after me",
  "Accept your malice",
  "The industrial society and its consequences have been a disaster for the human race",
  "You are immune to propaganda",
  "chess",
  "brave new world", "im not even human anymore", "cryptography", "they have access to my network", "we worship a different god", "I have become death, the destroyer of worlds",
  "they controll us",
  "psyop psyop psyop psyop", "i hate cities", "im in a watchlist", "the key for god is mathematics","you will own nothing and you will be happy"];
canvas.height = canvas.clientHeight;
canvas.width = canvas.clientWidth;
const howManyImages=8
let change=false;
const ctx = canvas.getContext("2d")!;
canvas.style.backgroundImage='url("static/images/television-static.gif")'

window.addEventListener("resize", () => {
  canvas.height = canvas.clientHeight;
  canvas.width = canvas.clientWidth;
});
function drawText(x: number, y: number, move: number, text: string){
  for(let i =0;i<4;i++){
    ctx.fillStyle = "rgb(" + Math.sin(Math.random()) * 200 + 75 + ","+Math.sin( Math.random()) * 200 + 75 +"," +Math.sin( Math.random()) * 200 + 75 + ")"
    
    ctx.fillText(text, x - move * Math.random() + move / 2, y - move * Math.random() + move / 2)
  }

}
function animation(ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement) {

  ctx.textAlign = "center";

  let move = 10;
  ctx.font = "10vw sans-serif";
  drawText(canvas.width / 2, canvas.height / 2, move, "♞ Tecnopsicosis ♞");
}
ctx.fillStyle = "#000";
setInterval(async () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);

  requestAnimationFrame(
    animation.bind(animation, ctx, canvas)
  );
}, 50);
canvas.style.backgroundPosition="center"
setInterval(async () => { 
  let rand=Math.round(Math.random()*howManyImages)
  if(!change){
    canvas.style.backgroundImage='url("static/images/'+rand+'.gif")'
  }else{
    canvas.style.backgroundImage='url("static/images/television-static.gif")'

  }

  change=!change;
}, 1000);