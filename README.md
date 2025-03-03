# piCRT

**piCRT** is a Raspberry Pi-powered retro media player written in Go and Svelte, with a terminal-style UI, designed for automated video playback and remote control.

---

## **📜 Features**

- 🔹 ASCII/BBS-style UI
- 🔹 Categorised media playback (e.g., Anime, JDM, Longplays)
- 🔹 Remote control via a web UI.
- 🔹 Runs on a Raspberry Pi with a Go backend & SvelteKit frontend.
- 🔹 Auto-starts on boot using `systemd`.

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
