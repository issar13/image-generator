const express = require("express");
const fs = require("fs");
const path = require("path");

const app = express();
const PORT = 3000;
const IMAGE_DIR = path.join(__dirname, "public/images");

app.use(express.static(path.join(__dirname, "public")));

// Serve image file names as JSON
app.get("/api/images", (req, res) => {
  fs.readdir(IMAGE_DIR, (err, files) => {
    if (err) {
      res.status(500).json({ error: "Unable to read image directory" });
    } else {
      const imageFiles = files.filter((file) =>
        /\.(jpe?g|png|gif|webp|bmp|heic)$/i.test(file)
      );
      res.json(imageFiles.map((file) => `images/${file}`));
    }
  });
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
