package chat

type inputChat struct {
	Text string `json:"text"`
}

type outputChat struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

const segmentAnswer string = `
Các bước cơ bản để tạo một User Segment mới:
1. Truy cập User segment>  Segmentation > Create new segment
2. Trong màn hình tạo User Segment, điền các thông tin: 
3. Chọn các điều kiện lọc phân khúc đối tượng theo yêu cầu. Chọn Thuộc tính người dùng hoặc Hành vi người dùng dựa trên yêu cầu của bạn.
4. Xem hướng dẫn chi tiết tại: https://cdp.gitbook.io/guide/phan-khuc-doi-tuong/tao-phan-khuc/chi-tiet-phan-khuc`

const ottAnswer string = `
Để có thể tạo 1 tin gửi OTT, bạn thực hiện các bước: 
1. Truy cập trang staff admin, chọn menu Tempate List -> chọn Create notification template
2. Chọn create OTT template
3. Thương hiệu: chọn một thương hiệu từ danh sách có sẵn
4. Tiêu đề là bắt buộc đối với tin thuộc loại OTT/ Email.
5. Xem hướng dẫn chi tiết tại: https://cdp.gitbook.io/guide/quan-ly-thong-bao/quan-ly-thong-bao`
