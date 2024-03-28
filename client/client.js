const HOST = `13.126.197.85`;
const PORT = 6000;
const LINK = "http://127.0.0.1:8080/channel?channelName=lessgo&password=deezBalls"
const { clear } = require("console");
const net = require("net");
const readline = require("readline/promises")
 
const rl = readline.createInterface({
    input : process.stdin,
    output : process.stdout
})

const clearline = (dir)=>{
    
    return new Promise ((resolve,reject)=>{
        process.stdout.clearLine(dir,()=>{
            resolve();
    });
});
}

const moveCursor = (dx,dy)=>{
    return new Promise((resolve ,reject)=>{
        process.stdout.moveCursor(dx,dy,()=>{
            resolve();
        })
    })    
}

let id;
console.log(`HOST  = ` + HOST);
 const client = net.createConnection({
    host:""
 },async ()=>{
   
    console.log("Connection created")

    const ask = async () =>{ 
        const message = await rl.question("Enter a message  > ");   
        
        if (message == '') ask();

        await moveCursor(0,-1)
        
        await clearline(0);//clears the current line cursor is currently in ..
        client.write(`${id}-message-${message}`);
     
    };
    
    ask();

    client.on("data",async (data)=>{

        console.log();

        await moveCursor(0,-1)
        
        await clearline(0);

        if (data.toString("utf-8").substring(0,3) === `id-`) {

            id = data.toString("utf-8").substring(3);
  
            console.log(" Your ID is " + id + "\n");
  
        }else {
  
            console.log(data.toString("utf-8"))
  
        }
  
        ask();
  
    })
})