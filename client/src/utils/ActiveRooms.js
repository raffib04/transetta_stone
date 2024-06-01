export const fetchActiveRooms = async () => {
    try {
        const res = await fetch(`${process.env.REACT_APP_BACKEND_URL}/rooms`)
        if (!res.ok) {
            throw new Error("Network response was not ok " + res.statusText);
        }
        const data = await res.json();
        return Array.isArray(data) ? data : [];
    } catch (error) {
        console.error("Error fetching rooms:", error);
        return [];
    }
};