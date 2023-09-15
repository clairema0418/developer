async function validateKey(key) {
  try {
    const response = await fetch(`/api/validate-key?key=${key}`);
    const data = await response.json();

    if (data.valid) {
      return { success: true, message: "Key is valid" };
    } else {
      return { success: false, message: "Key is invalid" };
    }
  } catch (error) {
    return { success: false, message: "Error validating key" };
  }
}

export default validateKey;