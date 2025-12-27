async function refreshAccessToken() {
    try {
      const res = await fetch("/refresh", {
        method: "POST",
        credentials: "include" 
      });
  
      if (res.ok) {
        console.log("Access token refreshed");
        return true;
      } else {
        console.log("Refresh token failed, user not logged in");
        return false;
      }
    } catch (err) {
      console.error("Error refreshing token:", err);
      return false;
    }
  }
  
  window.addEventListener("load", () => {
    refreshAccessToken();
  });
  