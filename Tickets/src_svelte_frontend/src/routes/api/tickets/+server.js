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

export async function POST({ url }) {
  const price = url.searchParams.get('price');
  const row = url.searchParams.get('row');
  const seat_number = url.searchParams.get('seat_number');
  const vip = url.searchParams.get('vip');
  const user_id = url.searchParams.get('user_id');
  const event_id = url.searchParams.get('event_id');

  const apiURL = `http://api:8000/tickets?price=${price}&row=${row}&seat_number=${seat_number}&vip=${vip}&user_id=${user_id}&event_id=${event_id}`;
  
  try {
    const response = await fetch(apiURL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        price: price,
        row: row,
        seat_number: seat_number,
        vip: vip,
        user_id: user_id,
        event_id: event_id
      })
    });

    const data = await response.json();
    if (response.ok) {
      return jsonResponse([data]);

    } else if (response.status == 400){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }

  } catch (error) {
    return jsonResponse({status: 500, error: 'Internal server error'});
  }
}