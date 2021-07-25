async function main() {
  let wasm;
  try{
    wasm = await fetch("particles.wasm");
  }catch(e) {
    console.log("error loading particles.wasm", e);
    return;
  }

  const go = new Go();
  let result;

  try {
    result = await WebAssembly.instantiateStreaming(wasm, go.importObject);
  } catch (e) {
    console.log("error instantiating webassembly", e);
    return;
  }

  go.run(result.instance);
}

main();
