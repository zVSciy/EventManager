function jsonResponse(json, status = 200) {
  return new Response(JSON.stringify(json), {
      status: status,
      headers: {
          'Content-Type': 'application/json'
      }
  });
}

// Bypass SSL certificate validation (development only)
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';
const BASE_DIR = 'https://tickets_api:8000';

export async function GET({ url }) {
  const userID = url.searchParams.get('user_id');
  const eventID = url.searchParams.get('event_id');
  let apiURL = '';
  
  if (!isNaN(eventID)) {
    apiURL = `${BASE_DIR}/tickets/user/${userID}?event_id=${eventID}`;
  } else {
    apiURL = `${BASE_DIR}/tickets/user/${userID}`;
  }
    
  try {
    const response = await fetch(apiURL);
    if (response.ok) {
      const data = await response.json();
      return jsonResponse(data);
    
    } else {
      return jsonResponse({status: response.status, error: data.detail.msg});
    }
  
  } catch (error) {
    return jsonResponse({status: 500, error: error});
  }
}
  
export async function DELETE({ url }) {
  const userID = url.searchParams.get('user_id');
  const ticketID = url.searchParams.get('ticket_id');
  
  if (isNaN(ticketID)) {
    return jsonResponse({status: 400, error: 'TicketID must be a number!'});
  }

  if (!isNaN(userID)) {
    return jsonResponse({status: 400, error: 'UserID must be a string!'});
  }

  const apiURL = `${BASE_DIR}/tickets/user/${userID}/ticket/${ticketID}`;

  try {
    const response = await fetch(apiURL, {
      method: 'DELETE'
    });

    const data = await response.json();
    if (response.ok) {
      return jsonResponse(data);

    } else {
      return jsonResponse({status: response.status, error: data.detail.msg});
    }
  } catch (error) {
    return jsonResponse({status: 500, error: error});
  }
}