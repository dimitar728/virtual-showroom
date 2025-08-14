# 🏛 Virtual Showroom & Booking System

A **3D interactive showroom platform** where users can explore virtual environments (car interiors, real estate rooms, trade show booths, etc.) and book in-person demos or tours.  
Admins can manage showroom content, booking slots, and user accounts through an intuitive dashboard.

---

## 📖 Table of Contents

-   Overview
    
-   Tech Stack
    
-   Features
    
-   System Architecture
    
-   Installation
    
-   API Endpoints
    
-   User Roles
    
-   Testing
    
-   Deployment
    
-   Screenshots / Demo
    
-   License
    

---

## 📌 Overview

The **Virtual Showroom & Booking System** allows businesses to:

-   Showcase 3D models of spaces or products in a web browser.
    
-   Offer customers an immersive first-person or orbit-controlled viewing experience.
    
-   Schedule and manage in-person demo appointments.
    
-   Maintain full admin control over showrooms, bookings, and users.
    

---

## 💻 Tech Stack

Layer

Technology

Frontend

React.js, Three.js

Backend

Go (Fiber or Gin)

Database

PostgreSQL

Auth

JWT

ORM

GORM

Storage

Local filesystem (S3-ready)

---

## ✨ Features

### 🖱 3D Showroom Explorer

-   Load `.glb`/`.gltf` models in an interactive Three.js viewer.
    
-   First-person or orbit navigation controls.
    
-   Hotspots with extra details.
    

### 👤 User Authentication & Booking

-   Secure signup/login with password hashing.
    
-   JWT-based authentication.
    
-   Book available time slots with calendar view.
    
-   Manage own bookings.
    

### 🛠 Admin Dashboard

-   Upload new showroom models with metadata.
    
-   Manage booking slots & blackout dates.
    
-   Suspend/reactivate/delete user accounts.
    

---

## 🏗 System Architecture

pgsql

CopyEdit

`[React + Three.js]  <-->  [Go API + GORM]  <-->  [PostgreSQL]                                  |                                Storage (Local / S3)`

---

## ⚙ Installation

### Prerequisites

-   Node.js ≥ 18
    
-   Go ≥ 1.20
    
-   PostgreSQL ≥ 14
    

### Backend Setup

bash

CopyEdit

`cd backendgo mod tidycp .env.example .env  # configure DB + JWT secretgo run main.go`

### Frontend Setup

bash

CopyEdit

`cd frontendnpm installcp .env.example .env  # configure API URLnpm start`

---

## 📡 API Endpoints (Sample)

Method

Endpoint

Description

Auth

POST

`/api/auth/register`

Register a new user

Public

POST

`/api/auth/login`

Login & get JWT

Public

GET

`/api/showrooms`

List all showrooms

User

POST

`/api/showrooms`

Create new showroom

Admin

POST

`/api/bookings`

Book a time slot

User

GET

`/api/admin/users`

List all users

Admin

---

## 👥 User Roles

Role

Capabilities

Visitor

View public showrooms, register, book slots

User

All visitor actions + manage own bookings

Admin

Full CRUD on showrooms, bookings, and users

---

## 🧪 Testing

-   **Unit Tests**: `go test ./...` for backend, `npm test` for frontend.
    
-   **API Tests**: Postman/Insomnia collections included in `/docs`.
    
-   **E2E Tests**: Cypress (optional).
    
-   **Performance Tests**: Locust / Artillery.
    

---

## ☁ Deployment

-   **Backend**: Render, Railway, Fly.io
    
-   **Frontend**: Vercel, Netlify, Firebase Hosting
    
-   **Database**: Supabase, Neon, Railway
    
-   **Model Storage**: AWS S3, Cloudinary
    

---

## 📸 Screenshots / Demo

*(Add screenshots or GIFs here)*

---

## 📜 License

MIT License — feel free to use and adapt.