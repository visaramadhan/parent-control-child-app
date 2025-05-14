# 👶 Child Care Apps for ADHD Child

A smart and caring mobile application designed to help parents monitor, manage, and support children with ADHD through device management, educational features, and real-time tracking tools.

---

## 📋 Description

**Child Care Apps for ADHD Child** is a mobile solution designed to help families and children manage ADHD-related challenges. It empowers parents to control screen time, block apps, track real-time locations, and promote educational rest time through interactive games.

---

## 🚀 Features

- 🔐 **Login System** – Secure login using Firebase Authentication  
- 🚫 **App Blocking** – Block specific apps to prevent distractions  
- 📵 **Device Control** – Remotely lock/unlock device functionality  
- ⏳ **Screen Time Monitoring** – Visualize and analyze usage data  
- 🧠 **Educational Games** – Stimulating games during rest periods  
- 📍 **Location Tracking** – Real-time GPS tracking for parents  
- 🔔 **Notification Alerts** – Get updates on child activity instantly  

---

## 🛠️ Technology Stack

### 📱 Frontend

- **Framework**: React Native (with Expo)
- **Navigation**: `@react-navigation/native`, `@react-navigation/bottom-tabs`
- **State Management**: Redux / React Context
- **API Handling**: Axios
- **UI Libraries**:
  - `react-native-paper` – Material Design components
  - `react-native-vector-icons` – Icon support
  - `lottie-react-native` – Animations
- **Charts**: `react-native-chart-kit` / `victory-native`
- **Firebase Integration**: 
  - `@react-native-firebase/auth`
  - `@react-native-firebase/messaging`
  - `@rn-bridge/react-native-geofencing`

#### 🎨 UI/UX Design

- **Design Tool**: Figma
- **Design System**: Material Design 3

---

### 🔧 Backend

- **Language**: Golang
- **Framework**: Gin
- **Routing**: Gorilla Mux
- **ORM**: GORM
- **JWT Auth**: `github.com/golang-jwt/jwt/v5`
- **Environment Management**: `github.com/joho/godotenv`
- **Firebase Admin SDK**: For messaging and real-time sync

#### 🗂️ API Modules

- `/auth`: Register/Login, JWT Token
- `/user`: Child-parent relationships, profile
- `/device`: App blocking, screen logs
- `/notification`: Push notifications via FCM
- `/location`: Geolocation sync and tracking

---

### 🔗 Database (via Firebase)

- **Authentication** – User login system
- **Firestore / RTDB** – Store logs, preferences, screen time
- **Firebase Cloud Messaging** – Real-time notification alerts
- **Firebase Geolocation** – Location updates in real-time

---

## ⚙️ Installation Guide

### 📁 Clone & Setup Project

```bash```
`git clone https://github.com/your-username/child-care-app.git`
`cd child-care-app`
`cp .env.example .env`

### 🤝 Contribution
- **Fork the project**
- **Create your branch**: `git checkout -b feature/YourFeature`
- **Commit changes**: `git commit -m 'Add YourFeature'`
- **Push branch**: `git push origin feature/YourFeature`

### 📧 Contact
Author: Visa Ramadhan
Email: visaramadhan28@gmail.com
GitHub: @visaramadhan

### 📄 License
This project is licensed under the MIT License.
See the LICENSE file for details.



