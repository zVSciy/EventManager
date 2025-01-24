document.getElementById('reviewForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const user_id = document.getElementById('user_id').value;
    const content = document.getElementById('content').value;
    const rating = document.getElementById('rating').value;
    const event_id = document.getElementById('event_id').value;

    if (user_id < 0 || event_id < 0) {
        alert('User ID and Event ID must be non-negative.');
        return;
    }

    const response = await fetch('http://ReviewAPI:8083/reviews/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ user_id, content, rating, event_id })
    });

    if (response.ok) {
        alert('Review submitted successfully!');
    } else {
        alert('Error submitting review.');
    }
});

document.getElementById('getReviewsForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const event_id = document.getElementById('event_id_get').value;

    if (event_id < 0) {
        alert('Event ID must be non-negative.');
        return;
    }

    const response = await fetch(`http://localhost:8083/reviews/${event_id}`);
    const reviews = await response.json();

    const reviewsDiv = document.getElementById('reviews');
    reviewsDiv.innerHTML = '<h2>Reviews:</h2>';
    reviews.forEach(review => {
        reviewsDiv.innerHTML += `<p>User ID: ${review.user_id}, Content: ${review.content}, Rating: ${review.rating}</p>`;
    });
});