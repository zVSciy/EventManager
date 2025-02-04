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
    apiURL = `http://tickets_api:8000/tickets?event_id=${eventID}`;
  } else {
    apiURL = `http://tickets_api:8000/tickets`;
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

export async function POST({ url }) {
  const price = url.searchParams.get('price');
  const row = url.searchParams.get('row');
  const seatNumber = url.searchParams.get('seat_number');
  const vip = url.searchParams.get('vip');
  const userID = url.searchParams.get('user_id');
  const eventID = url.searchParams.get('event_id');

  const apiURL = `http://tickets_api:8000/tickets`;
  let response = '';
  
  try {
    if (row == '' && seatNumber == '') {
      response = await fetch(apiURL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          price: price,
          vip: vip,
          user_id: userID,
          event_id: eventID
        })
      });

    } else if (row == '' || seatNumber == '') {
      return jsonResponse({status: 400, error: 'Row and seat number must be both filled or empty!'});
    
    } else {
      response = await fetch(apiURL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          price: price,
          row: row,
          seat_number: seatNumber,
          vip: vip,
          user_id: userID,
          event_id: eventID
        })
      });
    }

    const data = await response.json();
    if (response.ok) {
      return jsonResponse([data]);

    } else {
      return jsonResponse({status: response.status, error: data.detail.msg});
    }

  } catch (error) {
    return jsonResponse({status: 500, error: error});
  }
}

export async function PUT({ url }) {
  const changedTID = url.searchParams.get('ticket_id');
  const changedPrice = url.searchParams.get('price');
  const changedRow = url.searchParams.get('row');
  const changedSN = url.searchParams.get('seat_number');
  const changedVIP = url.searchParams.get('vip');
  const changedUID = url.searchParams.get('user_id');
  const changedEID = url.searchParams.get('event_id');

  if (isNaN(changedTID)) {
    return jsonResponse({status: 400, error: 'TicketID must be a number!'});
  }

  const apiURL = `http://tickets_api:8000/tickets/${changedTID}`;
  
  try {
    const response = await fetch(apiURL, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        price: changedPrice,
        row: changedRow,
        seat_number: changedSN,
        vip: changedVIP,
        user_id: changedUID,
        event_id: changedEID
      })
    });

    const data = await response.json();
    if (response.ok) {
      return jsonResponse([data]);

    } else {
      return jsonResponse({status: response.status, error: data.detail.msg});
    }

  } catch (error) {
    return jsonResponse({status: 500, error: error});
  }
}

export async function DELETE({ url }) {
  const ticketID = url.searchParams.get('ticket_id');
  
  if (isNaN(ticketID)) {
    return jsonResponse({status: 400, error: 'TicketID must be a number!'});
  }

  const apiURL = `http://tickets_api:8000/tickets/${ticketID}`;

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