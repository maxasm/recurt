const {encode} = require("gpt-3-encoder")

function count_tokens(content) {
	return encode(content).length 
} 

async function main() {
	process.stdin.setEncoding("utf-8")
	
	let all_text = ""
	
	process.stdin.on("readable", ()=> {
		let chunk;	
	
		while((chunk = process.stdin.read()) !== null) {
			all_text += chunk	
		} 
	})
	
	process.stdin.on("end", ()=> {
		let n_tokens = count_tokens(all_text)
		process.stdout.write(String(n_tokens))	
	})
}

main();
