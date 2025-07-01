# piCRT

**piCRT** is a Raspberry Pi-powered retro media player written in Go and Svelte, with a terminal-style UI, designed for automated video playback and remote control.

---

## **🛠️ Dev Mode & Media Path**

Want to use a different media folder for development or testing? Just set the `PICRT_MEDIA_PATH` environment variable when you run the server. This way, you don't have to edit the code to switch between your dev and production media folders.

**Why?**

- Makes it easy to test with a different set of videos on your laptop or dev machine.
- Keeps your production setup clean and safe.

**How to use:**

- **On Linux/macOS:**
  ```sh
  PICRT_MEDIA_PATH=/path/to/your/dev/media go run server.go
  # or if running the built binary
  PICRT_MEDIA_PATH=/path/to/your/dev/media ./server
  ```
- **On Windows (cmd):**
  ```cmd
  set PICRT_MEDIA_PATH=C:\path\to\your\dev\media
  go run server.go
  ```
- **On Windows (PowerShell):**
  ```powershell
  $env:PICRT_MEDIA_PATH="C:\path\to\your\dev\media"
  go run server.go
  ```
- **For production/systemd:**
  Add this to your service file:
  ```ini
  Environment=PICRT_MEDIA_PATH=/home/pi/media/
  ```

If you don't set it, the server defaults to `/home/pi/media/`.

---

## **📜 Features**

- 🔹 ASCII/BBS-style UI
- 🔹 Categorised media playback (e.g., Anime, JDM, Longplays)
- 🔹 Remote control via a web UI.
- 🔹 Runs on a Raspberry Pi with a Go backend & SvelteKit frontend.
- 🔹 Auto-starts on boot using `systemd`.

---

## Screenshots

![image](https://github.com/user-attachments/assets/408542aa-44d5-4b7c-bcb3-3685303133d2)

---

## **🚀 Installation & Setup**

### **1️⃣ Clone the Repository**

```bash
cd ~
git clone https://github.com/yourusername/piCRT.git
cd piCRT
```

### **2️⃣ Build & Deploy**

#### **On Your Dev Machine** (Arch/Linux/MacOS):

```bash
GOOS=linux GOARCH=arm64 go build -o server main.go
cd svelte-ui
npm install
npm run build
cd ..
```

#### **Transfer to Raspberry Pi**

```bash
scp -r server pi@<PI-IP>:/home/pi/piCRT/
scp -r svelte-ui/build pi@<PI-IP>:/home/pi/piCRT/build/
```

Or **SSH into your Raspberry Pi and use the deployment script:**

```bash
ssh pi@<PI-IP>
cd /home/pi/piCRT
./deploy-to-pi.sh
```

### **3️⃣ Setup `systemd` Service**

#### **On Raspberry Pi:**

```bash
sudo nano /etc/systemd/system/piCRT.service
```

**Paste the following:**

```ini
[Unit]
Description=piCRT Go Server
After=network.target

[Service]
ExecStart=/home/pi/piCRT/server
Restart=always
User=pi
WorkingDirectory=/home/pi/piCRT
StandardOutput=journal
StandardError=journal
Environment="PATH=/usr/local/bin:/usr/bin:/bin"
Environment=PICRT_MEDIA_PATH=/home/pi/media/

[Install]
WantedBy=multi-user.target
```

Save (`Ctrl+X`, `Y`, `Enter`), then enable & start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable piCRT
sudo systemctl start piCRT
```

### **4️⃣ Access the UI**

- Open `http://<PI-IP>:5000/` in a browser.
- Use the UI to browse & play videos.

---

## **🛠️ Development Workflow**

### **Updating Code & Deploying**

Use the provided script to **sync files and restart the server**:

```bash
./deploy-to-pi.sh
```

Or **SSH into your Raspberry Pi and run:**

```bash
ssh pi@<PI-IP>
cd /home/pi/piCRT
./deploy-to-pi.sh
```

### **Checking Logs & Debugging**

```bash
sudo systemctl status piCRT   # Check service status
journalctl -u piCRT --follow  # View logs
```

---

## **📜 TODO & Future Features**

- ✅ Add dynamic thumbnails for videos.
- ✅ Implement category filtering.
- ✅ Support for YouTube links.
- 🛠️ Improve mobile UI.

---

## **🖥️ Tech Stack**

- **Backend:** Go
- **Frontend:** SvelteKit + Tailwind CSS
- **Deployment:** `systemd` + SSH + SCP
