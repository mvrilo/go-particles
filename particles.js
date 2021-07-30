async function main() {
  if (!WebAssembly.instantiateStreaming) {
    // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }

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
