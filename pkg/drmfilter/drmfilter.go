package drmfilter

func Filter(req *Request) *Response {
	resp := &Response{}
	for _, r := range req.Payload {
		if r.Drm && r.EpisodeCount > 0 {
			item := ResponseItem{Image: r.Image.ShowImage,
				Title: r.Title,
				Slug:  r.Slug}
			resp.Response = append(resp.Response, item)
		}
	}
	return resp
}
