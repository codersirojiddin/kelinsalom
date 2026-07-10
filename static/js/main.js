// Sahifa yuklanganda saqlangan shrift o'lchamini tikshirish
document.addEventListener("DOMContentLoaded", function() {
    const savedSize = localStorage.getItem("userFontSize");
    if (savedSize) {
        document.body.style.fontSize = savedSize + "px";
    }

    // Dinamik qidiruvni faollashtirish
    initLiveSearch();
});

// Shriftni o'zgartirish funksiyasi
function changeFontSize(step) {
    const currentSize = parseFloat(window.getComputedStyle(document.body, null).getPropertyValue('font-size'));
    let newSize = currentSize + step;
    if (newSize < 16) newSize = 16;
    if (newSize > 32) newSize = 32;
    document.body.style.fontSize = newSize + "px";
    localStorage.setItem("userFontSize", newSize);
}

// === LIVE SEARCH & DEBOUNCE MANTIQI ===
function initLiveSearch() {
    const searchInput = document.getElementById("searchInput");
    const regionSelect = document.getElementById("regionSelect");
    
    if (!searchInput || !regionSelect) return;

    let debounceTimer;

    // Qidiruv inputiga yozganda
    searchInput.addEventListener("input", function() {
        clearTimeout(debounceTimer);
        // Foydalanuvchi yozishdan to'xtagach 500ms (0.5 soniya) kutadi
        debounceTimer = setTimeout(() => {
            fetchFilteredPoems();
        }, 500);
    });

    // Viloyat o'zgarganda srazi qidiradi (kutib o'tirmaydi)
    regionSelect.addEventListener("change", function() {
        fetchFilteredPoems();
    });
}

// Sahifani reload qilmasdan backend'dan ma'lumotni Fetch orqali olish
function fetchFilteredPoems() {
    const query = document.getElementById("searchInput").value;
    const region = document.getElementById("regionSelect").value;
    const poemsList = document.querySelector(".poems-list");

    // URL'ni shakllantirish
    const url = `/?q=${encodeURIComponent(query)}&region=${encodeURIComponent(region)}`;

    // Orqa fonda HTML so'rab olish
    fetch(url)
        .then(response => response.text())
        .then(html => {
            // Kelgan HTML ichidan faqat ".poems-list" maydonini qirqib olib sahifaga joylash
            const parser = new DOMParser();
            const doc = parser.parseFromString(html, 'text/html');
            const newPoemsList = doc.querySelector(".poems-list");
            
            if (newPoemsList && poemsList) {
                poemsList.innerHTML = newPoemsList.innerHTML;
            }
        })
        .catch(err => console.error("Qidiruvda xatolik:", err));
}