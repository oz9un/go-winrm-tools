
# WinRM Tools

This repo includes several **WinRM** tools written with **Go**:

1. File transfering between two Powershell session.
2. Running command on remote Powershell session.

#### I used [masterzen](https://github.com/masterzen/winrm)'s amazing repository for WinRM connections.


## Feature 1 - File Transfer 🛵

You can **transfer**/**receive** files **to**/**from** a Powershell session with an interactive screen.

![file-transfer](https://user-images.githubusercontent.com/57866851/143190832-92904a3d-a8b1-47e6-b0f9-289014cbcce6.png)

### Parameters:
File transfer tool takes 3 parameters:

1. **Direction:** It denotes the direction of the file.
- **1** for from *local* to *remote*.
- **2** for from *remote* to *local*.
2. **First File:** File's path that you want to send or receive.
3. **Second File:** File's target path.


## Feature 2 - Command Runner 💂

You can run commands on **remote *Powershell* session** with an interactive screen.

![command-runner](https://user-images.githubusercontent.com/57866851/143191697-c2ce9cb8-e26a-41b6-ad7e-b56412435825.png)

### Parameters:

Command runner tool takes no parameters, you can use it directly from the interactive screen.
