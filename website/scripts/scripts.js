/**
 * Submits a form to shorten a given URL. It retrieves the value in the #long-url
 * input box, sends a POST request to the /shorten endpoint, and then updates the
 * #short-url <p> element with the shortened URL.
 */
async function submitForm() {
  const url = document.getElementById("long-url").value;
  console.log("url:", url);
  const fetchUrl = async () => {
    const response = await fetch("/shorten", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url }),
    });
    const data = await response.json();
    return data;
  };

  data = await fetchUrl();
  console.log("data:", data);
  document.getElementById("short-url").textContent = data.short_url;
}
