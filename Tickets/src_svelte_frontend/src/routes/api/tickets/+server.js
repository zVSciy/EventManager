function jsonResponse(json, status = 200) {
  return new Response(JSON.stringify(json), {
      status: status,
      headers: {
          'Content-Type': 'application/json'
      }
  });
}

export async function GET({ url }) {
  const eventID = url.searchParams.get('event_id');
  let apiURL = '';

  if (!isNaN(eventID)) {
    apiURL = `http://api:8000/tickets?event_id=${eventID}`;
  } else {
    apiURL = `http://api:8000/tickets`;
  }
  
  try {
    const response = await fetch(apiURL);
    if (response.ok) {
      const data = await response.json();
      return jsonResponse(data);
    
    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }

  } catch (error) {
    return jsonResponse({status: 500, error: 'Internal server error'});
  }
}