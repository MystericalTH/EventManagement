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
	http.delete('/api/members/requests/:id', ({ params }) => {
		return new HttpResponse(null, { status: 204 });
	}),
	http.post('/api/members/requests/:id/approve', ({ params }) => {
		return new HttpResponse(null, { status: 204 });
	}),
	http.get('/api/activities', () => {
		return HttpResponse.json(activityData);
	}),
	http.get('/api/activities/:id', ({ params }) => {
		let { id } = params;
		console.log('enter');
		return HttpResponse.json(activityData.at(id));
	}),
	http.get('/api/activities/:id/registration/status', ({}) => {
		return HttpResponse.json({ is_registered: true });
	}),
	http.get('/api/activities/:id/feedback/status', ({}) => {
		return HttpResponse.json({ hasSubmittedFeedback: true });
	})
];
