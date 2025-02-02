document.getElementById('reviewForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const user_id = document.getElementById('user_id').value;
    const content = document.getElementById('content').value;
    const rating = document.getElementById('rating').value;
    const event_id = document.getElementById('event_id').value;

    console.log('Submitting review:', { user_id, content, rating, event_id });

    if (user_id < 0 || event_id < 0) {
        alert('User ID and Event ID must be non-negative.');
        console.warn('User ID or Event ID is negative:', { user_id, event_id });
        return;
    }

    try {
        const response = await fetch('http://ReviewAPI:8083/reviews/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ user_id, content, rating, event_id })
        });

        if (response.ok) {
            alert('Review submitted successfully!');
            console.log('Review submitted successfully:', { user_id, event_id });
        } else {
            alert('Error submitting review.');
            console.error('Error submitting review:', response.statusText);
        }
    } catch (error) {
        alert('Error submitting review.');
        console.error('Error submitting review:', error);
    }
});

document.getElementById('getReviewsForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const event_id = document.getElementById('event_id_get').value;

    console.log('Fetching reviews for event:', { event_id });

    if (event_id < 0) {
        alert('Event ID must be non-negative.');
        console.warn('Event ID is negative:', { event_id });
        return;
    }

    try {
        const response = await fetch(`http://192.168.75.138:8083/reviews/event/${event_id}`);
        const reviews = await response.json();

        console.log('Fetched reviews:', reviews);

        const reviewsDiv = document.getElementById('reviews');
        reviewsDiv.innerHTML = '<h2>Reviews:</h2>';
        reviews.forEach(review => {
            reviewsDiv.innerHTML += `<pre>${JSON.stringify(review, null, 2)}</pre>`;
        });
    } catch (error) {
        alert('Error fetching reviews.');
        console.error('Error fetching reviews:', error);
    }
});

document.getElementById('updateReviewForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const review_id = document.getElementById('review_id_update').value;
    const content = document.getElementById('content_update').value;
    const rating = document.getElementById('rating_update').value;

    console.log('Updating review:', { review_id, content, rating });

    try {
        const response = await fetch(`http://192.168.75.138:8083/reviews/${review_id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ content, rating })
        });

        if (response.ok) {
            alert('Review updated successfully!');
            console.log('Review updated successfully:', { review_id });
        } else {
            alert('Error updating review.');
            console.error('Error updating review:', response.statusText);
        }
    } catch (error) {
        alert('Error updating review.');
        console.error('Error updating review:', error);
    }
});

document.getElementById('deleteReviewForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const review_id = document.getElementById('review_id_delete').value;

    console.log('Deleting review:', { review_id });

    try {
        const response = await fetch(`http://192.168.75.138:8083/reviews/${review_id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            alert('Review deleted successfully!');
            console.log('Review deleted successfully:', { review_id });
        } else {
            alert('Error deleting review.');
            console.error('Error deleting review:', response.statusText);
        }
    } catch (error) {
        alert('Error deleting review.');
        console.error('Error deleting review:', error);
    }
});

document.getElementById('getReviewForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const review_id = document.getElementById('review_id_get').value;

    console.log('Fetching review by ID:', { review_id });

    try {
        const response = await fetch(`http://192.168.75.138:8083/reviews/${review_id}`);
        const review = await response.json();

        console.log('Fetched review:', review);

        const reviewsDiv = document.getElementById('reviews');
        reviewsDiv.innerHTML = '<h2>Review:</h2>';
        reviewsDiv.innerHTML += `<pre>${JSON.stringify(review, null, 2)}</pre>`;
    } catch (error) {
        alert('Error fetching review.');
        console.error('Error fetching review:', error);
    }
});