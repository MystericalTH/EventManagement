import { http } from 'msw';
import { HttpResponse } from 'msw';
import { memberData, activityData } from './__data';
export const handlers = [
	// Mock a GET request to "/api/users"
	http.get('/api/users', ({ params }) => {
		return HttpResponse.json([
			{
				name: 'Nattawin Yamprasert',
				email: 'test@email.com'
			},
			{
				name: 'EIEIE YAYY',
				email: 'test2@email.com'
			}
		]);
	}),

	http.get('/api/members/requests', ({ params }) => {
		return HttpResponse.json(memberData);
	}),
	http.get('/api/members', ({}) => {
		return HttpResponse.json(memberData);
	}),
	http.delete('/api/members/:id', ({ params }) => {
		return new HttpResponse(null, { status: 204 });
	}),
	http.put('/api/members/requests/:id/approve', ({ params }) => {
		return new HttpResponse(null, { status: 204 });
	}),
	http.get('/api/activities/requests', () => {
		return HttpResponse.json(activityData);
	}),
	http.get('/api/activities', () => {
		return HttpResponse.json(activityData);
	}),
	http.put('/api/activities/:id/approve', ({ params }) => {
		return new HttpResponse(null, { status: 204 });
	}),
	http.delete('/api/activities/:id', ({ params }) => {
		return new HttpResponse(null, { status: 204 });
	})
];
