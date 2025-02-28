function jsonResponse(json, status = 200) {
  return new Response(JSON.stringify(json), {
      status: status,
      headers: {
          'Content-Type': 'application/json'
      }
  });
}

export async function PUT({ url }) {
  const eventID = url.searchParams.get('event_id');
  const moreTickets = url.searchParams.get('delete');
  const vip = url.searchParams.get('vip');
  let normal_tickets = 2;
  let vip_tickets = 2;
  console.log(vip)
  console.log(moreTickets)
  if (vip == 'true') {
  
    if (moreTickets == 0) {
      vip_tickets = 0;
    }else if (moreTickets == 1) {
      vip_tickets = 1;
    }
    
  } else {
  
    if (moreTickets == 1) {
      normal_tickets = 1;}
      else if (moreTickets == 0) {
        normal_tickets = 0;
      }
  
  }
  if (isNaN(eventID)) {
    return jsonResponse({status: 400, error: 'EventID must be a number!'});
  }

  const apiURL = `http://events_api:8000/event/updateTicket/${eventID}`;
  
  try {
    const response = await fetch(apiURL, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        available_normal_tickets: normal_tickets ,
        available_vip_tickets: vip_tickets
      })
    });

    const data = await response.json();
    console.log(data)
    if (response.ok) {
      return jsonResponse([data]);

    } else if (response.status == 400){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else if (response.status == 404){
      return jsonResponse({status: response.status, error: data.detail.msg});

    } else {
      return jsonResponse({status: response.status, error: 'Failed to fetch data!'});
    }

  } catch (error) {
    return jsonResponse({status: 500, error: error});
  }
}
