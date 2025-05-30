const fs = require("fs");
const path = require("path");
const heicConvert = require("heic-convert");

const inputDir = "./images"; // or whatever path you're using

function getOutputPath(filePath) {
  const ext = path.extname(filePath);
  const base = path.basename(filePath, ext);
  return path.join(inputDir, `${base}.jpg`);
}

async function convertHeicToJpg(filePath) {
  const outputPath = getOutputPath(filePath);

  if (fs.existsSync(outputPath)) {
    console.log(`✅ Skipped (already exists): ${path.basename(outputPath)}`);
    return;
  }

  try {
    const inputBuffer = fs.readFileSync(filePath);
    const outputBuffer = await heicConvert({
      buffer: inputBuffer,
      format: "JPEG",
      quality: 0.8,
    });

    fs.writeFileSync(outputPath, outputBuffer);
    console.log(
      `🎉 Converted: ${path.basename(filePath)} → ${path.basename(outputPath)}`
    );
  } catch (err) {
    console.error(`❌ Failed: ${filePath}`, err.message);
  }
}

async function run() {
  console.log(`📁 Scanning folder: ${path.resolve(inputDir)}`);
  const files = fs.readdirSync(inputDir);

  const heicFiles = files.filter((file) =>
    file.toLowerCase().endsWith(".heic")
  );

  console.log(`📸 Found ${heicFiles.length} HEIC file(s).`);

  for (const file of heicFiles) {
    await convertHeicToJpg(path.join(inputDir, file));
  }

  if (heicFiles.length === 0) {
    console.log("⚠️ No HEIC files found.");
  }
}

run();
