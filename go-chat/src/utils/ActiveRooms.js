export const fetchActiveRooms = async () => {
    try {
        const res = await fetch("http://localhost:8080/rooms");
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