const fs = require("fs");
const path = require("path");
const sharp = require("sharp");

const inputDir = "./public/images";
const outputDir = path.join(inputDir, "compressed-webp");

// Ensure output directory exists
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir, { recursive: true });
}

function isImageFile(fileName) {
  const ext = path.extname(fileName).toLowerCase();
  return [".jpg", ".jpeg", ".png"].includes(ext);
}

async function compressToWebp(inputFile, outputFile) {
  try {
    const stats = fs.statSync(inputFile);
    if (stats.size < 100 * 1024) {
      console.log(`⚠️ Skipped (already small): ${path.basename(inputFile)}`);
      return;
    }

    await sharp(inputFile)
      .resize({ width: 1280, height: 1280, fit: "inside" }) // Optional resizing
      .webp({ quality: 60 }) // Reduce quality to shrink further
      .toFile(outputFile);

    const newSize = fs.statSync(outputFile).size / 1024;
    console.log(
      `✅ Compressed: ${path.basename(inputFile)} → ${Math.round(newSize)} KB`
    );
  } catch (err) {
    console.error(`❌ Failed to compress ${inputFile}`, err.message);
  }
}

async function run() {
  console.log(`📁 Scanning folder: ${path.resolve(inputDir)}`);
  const files = fs.readdirSync(inputDir);

  const imageFiles = files.filter(isImageFile);
  console.log(`🖼️ Found ${imageFiles.length} image(s) to compress.`);

  for (const file of imageFiles) {
    const inputFile = path.join(inputDir, file);
    const outputFile = path.join(outputDir, path.parse(file).name + ".webp");
    await compressToWebp(inputFile, outputFile);
  }

  if (imageFiles.length === 0) {
    console.log("⚠️ No image files found.");
  }
}

run();
