package services

import (
	"admin_svc/pkg/db"
	"admin_svc/pkg/models"
	pb "admin_svc/pkg/pb"
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type Server struct {
	H db.Handler
	pb.UnimplementedAdminServiceServer
}

func (s *Server) Feeds(ctx context.Context, req *pb.FeedsRequest) (*pb.FeedsResponse, error) {

	log.Println("Feeds collection started")
	log.Println("Data collected", req)
	var page, limit int64
	page, limit = int64(req.Page), int64(req.Limit)
	// pagination purpose -
	if req.Page == 0 {
		page = 1
	}
	if req.Limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var postdetails []*pb.Post

	sqlQuery := "SELECT * FROM posts WHERE status = 'approved'"
	if req.Searchkey != "" {
		sqlQuery += " AND (text ILIKE '%" + req.Searchkey + "%' OR place ILIKE '%" + req.Searchkey + "%')"
	}
	sqlQuery += " ORDER BY date DESC, amount DESC LIMIT ? OFFSET ?"

	if err := s.H.DB.Raw(sqlQuery, limit, offset).Scan(&postdetails).Error; err != nil {
		return &pb.FeedsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
			Posts:    []*pb.Post{},
		}, err
	}
	log.Println("feeds:", postdetails, req, limit, page)
	return &pb.FeedsResponse{
		Status:   http.StatusOK,
		Response: "success",
		Posts:    postdetails,
	}, nil

}

func (s *Server) PostDetails(ctx context.Context, req *pb.PostDetailsRequest) (*pb.PostDetailsResponse, error) {

	log.Println("Post detailes started")

	var postdetails *pb.Post

	if err := s.H.DB.Raw("SELECT * FROM posts where id=?", req.PostID).Scan(&postdetails).Error; err != nil {
		return &pb.PostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("post:", postdetails)
	return &pb.PostDetailsResponse{
		Status:   http.StatusOK,
		Response: "success",
		Post:     postdetails,
	}, nil

}

func (s *Server) ApproveCampaign(ctx context.Context, req *pb.ApproveCampaignRequest) (*pb.ApproveCampaignResponse, error) {

	log.Println("Approve Campaign started")

	var postdetails *pb.Post

	err := s.H.DB.Exec("UPDATE posts set status = 'approved' where id = ?", req.Id).Error
	if err != nil {
		log.Println(err)
		return &pb.ApproveCampaignResponse{Status: http.StatusBadGateway, Response: err.Error()}, err
	}

	if err := s.H.DB.Raw("SELECT * FROM posts where id=?", req.Id).Scan(&postdetails).Error; err != nil {
		return &pb.ApproveCampaignResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("post:", postdetails)
	return &pb.ApproveCampaignResponse{
		Status:   http.StatusOK,
		Response: "success",
		//Post:   postdetails,
	}, nil

}

func (s *Server) RejectCampaign(ctx context.Context, req *pb.RejectCampaignRequest) (*pb.RejectCampaignResponse, error) {

	log.Println("Reject Campaign started")

	var postdetails *pb.Post

	err := s.H.DB.Exec("UPDATE posts set status = 'rejected' where id = ?", req.Id).Error
	if err != nil {
		log.Println(err)
		return &pb.RejectCampaignResponse{Status: http.StatusBadGateway, Response: err.Error()}, err
	}

	if err := s.H.DB.Raw("SELECT * FROM posts where id=?", req.Id).Scan(&postdetails).Error; err != nil {
		return &pb.RejectCampaignResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("post:", postdetails)
	return &pb.RejectCampaignResponse{
		Status:   http.StatusOK,
		Response: "success",
		//Post:   postdetails,
	}, nil

}

func (s *Server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {

	log.Println("Delete Post started")

	err := s.H.DB.Exec("delete from posts where id = ?", req.Id).Error
	if err != nil {
		log.Println(err)
		return &pb.DeletePostResponse{Status: http.StatusBadGateway, Response: err.Error()}, err
	}

	return &pb.DeletePostResponse{
		Status:   http.StatusOK,
		Response: "success",
	}, nil

}

func (s *Server) CampaignDetails(ctx context.Context, req *pb.CampaignDetailsRequest) (*pb.CampaignDetailsResponse, error) {

	log.Println("Campaign detailes started")

	var CampaignDetails *pb.Post

	if err := s.H.DB.Raw("SELECT * FROM posts where id=?", req.Id).Scan(&CampaignDetails).Error; err != nil {
		return &pb.CampaignDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("campaign:", CampaignDetails)
	return &pb.CampaignDetailsResponse{
		Status:   http.StatusOK,
		Response: "success",
		Post:     CampaignDetails,
	}, nil

}

func (s *Server) CampaignRequestList(ctx context.Context, req *pb.CampaignRequestListRequest) (*pb.CampaignRequestListResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	var page, limit int64
	page, limit = int64(req.Page), int64(req.Limit)
	// pagination purpose -
	if req.Page == 0 {
		page = 1
	}
	if req.Limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var postdetails []*pb.Post

	sqlQuery := "SELECT * FROM posts WHERE status = 'pending'"
	if req.Searchkey != "" {
		sqlQuery += " AND (text ILIKE '%" + req.Searchkey + "%' OR place ILIKE '%" + req.Searchkey + "%')"
	}
	sqlQuery += " ORDER BY date DESC, amount DESC LIMIT ? OFFSET ?"
log.Println(" Query:",sqlQuery)
	if err := s.H.DB.Raw(sqlQuery, limit, offset).Scan(&postdetails).Error; err != nil {
		return &pb.CampaignRequestListResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
			Post:     []*pb.Post{},
		}, err
	}
	log.Println("feeds:", postdetails, req, limit, page)
	return &pb.CampaignRequestListResponse{
		Status:   http.StatusOK,
		Response: "successfully colected list",
		Post:     postdetails,
	}, nil

}

// //////////////////////////////////////////////////////////////////////////////////////////
func (s *Server) ReportedList(ctx context.Context, req *pb.ReportedListRequest) (*pb.ReportedListResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	var page, limit int64
	page, limit = int64(req.Page), int64(req.Limit)
	// pagination purpose -
	if req.Page <= 0 {
		page = 1
	}
	if req.Limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var reportedposts []*models.Reported
	sqlQuery := "SELECT * FROM reporteds"
	if req.Searchkey != "" {
		sqlQuery += " WHERE reason ILIKE '%" + req.Searchkey + "%'"
	}
	sqlQuery += " ORDER BY id DESC LIMIT ? OFFSET ?"

	if err := s.H.DB.Raw(sqlQuery, limit, offset).Scan(&reportedposts).Error; err != nil {
		return &pb.ReportedListResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("reports: ",reportedposts)
	var reporteds []*pb.ReportedPost
	for _, v := range reportedposts {
		post := new(pb.ReportedPost)
		post.Reason = v.Reason
		if err := s.H.DB.Raw("SELECT * FROM posts where id=?", v.PostID).Scan(&post.Post).Error; err != nil {
			log.Printf("Failed to retrieve post details for reported post %d: %v", v.PostID, err)
			continue
		}
		reporteds = append(reporteds, post)
	}
	log.Println("Data Recieved: ", req)
	return &pb.ReportedListResponse{
		Status:   http.StatusOK,
		Response: "Successfully collected reported posts",
		Post:     reporteds,
	}, nil

}
func (s *Server) ReportDetails(ctx context.Context, req *pb.ReportDetailsRequest) (*pb.ReportDetailsResponse, error) {
	log.Println("Campaign Request List started")
	log.Println("Data collected", req)

	var reportedPost models.Reported
	sqlQuery := "SELECT * FROM reporteds where id=?"
	if err := s.H.DB.Raw(sqlQuery, req.Postid).Scan(&reportedPost).Error; err != nil {
		return &pb.ReportDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}

	var post pb.Post
	if err := s.H.DB.Raw("SELECT * FROM posts WHERE id=?", req.Postid).Scan(&post).Error; err != nil {
		log.Printf("Failed to retrieve post details: %v", err)
		return &pb.ReportDetailsResponse{
			Status:   http.StatusInternalServerError,
			Response: "Failed to retrieve post details",
		}, err
	}

	postDetails := &pb.ReportedPost{
		Reason: reportedPost.Reason,
		Post:   &post,
	}
	log.Println("Data Recieved: ", req)
	return &pb.ReportDetailsResponse{
		Status:   http.StatusOK,
		Response: "Successfully fetched the Reported post",
		Post:     postDetails,
	}, nil

}
func (s *Server) DeleteReport(ctx context.Context, req *pb.DeleteReportRequest) (*pb.DeleteReportResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	err := s.H.DB.Exec("delete from reporteds where id = ?", req.Postid).Error
	if err != nil {
		return &pb.DeleteReportResponse{Status: http.StatusBadGateway, Response: "Couldnt delete the report"}, err
	}
	log.Println("Data Recieved: ", req)
	return &pb.DeleteReportResponse{
		Status:   http.StatusOK,
		Response: "Successfully deleted report",
	}, nil

}
func (s *Server) CategoryList(ctx context.Context, req *pb.CategoryListRequest) (*pb.CategoryListResponse, error) {

	log.Println("Category Request List started")
	log.Println("Data collected", req)
	var page, limit int64
	page, limit = int64(req.Page), int64(req.Limit)
	// pagination purpose -
	if req.Page <= 0 {
		page = 1
	}
	if req.Limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var categoryList []*pb.Category
	sqlQuery := "SELECT * FROM categories"
	if req.Searchkey != "" {
		sqlQuery += " WHERE reason ILIKE '%" + req.Searchkey + "%'"
	}
	sqlQuery += " ORDER BY id DESC LIMIT ? OFFSET ?"

	if err := s.H.DB.Raw(sqlQuery, limit, offset).Scan(&categoryList).Error; err != nil {
		return &pb.CategoryListResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("Data Recieved: ", req)
	return &pb.CategoryListResponse{
		Status:   http.StatusOK,
		Response: "Successfully Fetched the data",
		Categories: categoryList,
	}, nil

}
func (s *Server) CategoryPosts(ctx context.Context, req *pb.CategoryPostsRequest) (*pb.CategoryPostsResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	var page, limit int64
	page, limit = int64(req.Page), int64(req.Limit)
	// pagination purpose -
	if req.Page <= 0 {
		page = 1
	}
	if req.Limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	if req.Categoryid < 1 {
		return &pb.CategoryPostsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid category ID",
		}, errors.New("invalid category ID")
	}
	//get category name
	var category *pb.Category
	sqlQuery := "SELECT * FROM categories where id=?"
	if err := s.H.DB.Raw(sqlQuery, req.Categoryid).Scan(&category).Error; err != nil {
		return &pb.CategoryPostsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get categories from DB",
			Category: category,
		}, err
	}
	//now get posts
	var post []*pb.Post

	sqlQuery = "SELECT * FROM posts where cat_id=?'"
	sqlQuery += " ORDER BY id DESC LIMIT ? OFFSET ?"
	if err := s.H.DB.Raw(sqlQuery, req.Categoryid, limit, offset).Scan(&category).Error; err != nil {
		return &pb.CategoryPostsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get categories from DB",
			Category: category,
		}, err
	}
	log.Println("Data Recieved: ", req)
	return &pb.CategoryPostsResponse{
		Status:   http.StatusOK,
		Response: "fetched the Category posts",
		Category: category,
		Posts:    post,
	}, nil

}
func (s *Server) NewCategory(ctx context.Context, req *pb.NewCategoryRequest) (*pb.NewCategoryResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	query := `
	INSERT INTO categories (category)
	VALUES (?);
`
	if err := s.H.DB.Exec(query, req.Category).Error; err != nil {
		log.Printf("Failed to create category: %v", err)
		return &pb.NewCategoryResponse{
			Status:   http.StatusInternalServerError,
			Response: "Failed to create category",
		}, err
	}
	log.Println("Data Recieved: ", req)
	return &pb.NewCategoryResponse{
		Status:   http.StatusOK,
		Response: "Successfully created new category",
		Category: &pb.Category{Category: req.Category},
	}, nil

}
func (s *Server) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	err := s.H.DB.Exec("delete from categories where id = ?", req.Categoryid).Error
	if err != nil {
		return &pb.DeleteCategoryResponse{Status: http.StatusBadGateway, Response: "Couldnt delete the category"}, err
	}
	log.Println("Data Recieved: ", req)
	return &pb.DeleteCategoryResponse{
		Status:   http.StatusOK,
		Response: "Category Deleted Successfully",
	}, nil

}
func (s *Server) AdminDashboard(ctx context.Context, req *pb.AdminDashboardRequest) (*pb.AdminDashboardResponse, error) {

	now := time.Now()

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	startOfWeek := now.AddDate(0, 0, -int(now.Weekday())+1)

	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())

	daily, err := s.GetStats(startOfDay)
	if err != nil {
		return &pb.AdminDashboardResponse{Status: http.StatusBadGateway, Response: "Couldnt fetch the stats"}, err
	}
	weekly, err := s.GetStats(startOfWeek)
	if err != nil {
		return &pb.AdminDashboardResponse{Status: http.StatusBadGateway, Response: "Couldnt fetch the stats"}, err
	}
	monthly, err := s.GetStats(startOfMonth)
	if err != nil {
		return &pb.AdminDashboardResponse{Status: http.StatusBadGateway, Response: "Couldnt fetch the stats"}, err
	}
	annual, err := s.GetStats(startOfYear)
	if err != nil {
		return &pb.AdminDashboardResponse{Status: http.StatusBadGateway, Response: "Couldnt fetch the stats"}, err
	}
	alltime, err := s.GetStats(startOfYear)
	if err != nil {
		return &pb.AdminDashboardResponse{Status: http.StatusBadGateway, Response: "Couldnt fetch the stats"}, err
	}

	log.Println("Data Received: ", req)
	return &pb.AdminDashboardResponse{
		Status:   http.StatusOK,
		Response: "Successfully Fetched Stats",
		Daily:    daily,
		Weekly:   weekly,
		Monthly:  monthly,
		Annual:   annual,
		Alltime:  alltime,
	}, nil

}
func (s *Server) PostStats(ctx context.Context, req *pb.PostStatsRequest) (*pb.PostStatsResponse, error) {

	log.Println("Campaign Request List started")
	log.Println("Data collected", req)
	var page, limit int64
	page, limit = int64(req.Page), int64(req.Limit)
	// pagination purpose -
	if req.Page <= 0 {
		page = 1
	}
	if req.Limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var posts []*pb.Post
	sqlQuery := "SELECT * FROM posts"
	sqlQuery += " ORDER BY views DESC,likes DESC LIMIT ? OFFSET ?"

	if err := s.H.DB.Raw(sqlQuery, limit, offset).Scan(&posts).Error; err != nil {
		return &pb.PostStatsResponse{
			Status:   http.StatusBadRequest,
			Response: "couldn't get posts from DB",
		}, err
	}
	log.Println("Data Recieved: ", req)
	return &pb.PostStatsResponse{
		Status:   http.StatusOK,
		Response: "Successfully Got Trending list",
		Posts:    posts,
	}, nil

}

// func (s *Server) getStats(date time.Time) *pb.Stats {
// 	collectedMoney := 0
// 	sqlQuery := "SELECT SUM(collected) as collected_money FROM posts where date > ?"
// 	if err := s.H.DB.Raw(sqlQuery, date).Scan(&collectedMoney).Error; err != nil {
// 		return &pb.Stats{}
// 	}

// 	posts := 0
// 	sqlQuery = "SELECT COUNT(*) as posts FROM posts where date > ? and (status = 'approved' or status = 'expired'"
// 	if err := s.H.DB.Raw(sqlQuery, date).Scan(&posts).Error; err != nil {
// 		return &pb.Stats{Collectedmoney: int64(collectedMoney)}
// 	}
// 	likes := 0
// 	sqlQuery = "SELECT SUM(likes) as likes FROM posts where date > ?"
// 	if err := s.H.DB.Raw(sqlQuery, date).Scan(&likes).Error; err != nil {
// 		return &pb.Stats{Collectedmoney: int64(collectedMoney),
// 			Posts: int64(posts)}
// 	}
// 	campaigns := 0
// 	sqlQuery = "SELECT COUNT(*) as campaigns FROM posts where date > ? and status = 'approved'"
// 	if err := s.H.DB.Raw(sqlQuery, date).Scan(&campaigns).Error; err != nil {
// 		return &pb.Stats{Collectedmoney: int64(collectedMoney),
// 			Posts: int64(posts),
// 			Likes: int64(likes)}
// 	}
// 	return &pb.Stats{
// 		Collectedmoney:  int64(collectedMoney),
// 		Posts:           int64(posts),
// 		Likes:           int64(likes),
// 		Activecampaigns: int64(campaigns),
// 	}
// }
func (s *Server) GetStats(date time.Time) (*pb.Stats, error) {
	var stats pb.Stats

	tx := s.H.DB.Begin()
	defer tx.Rollback() // Rollback the transaction if it's not committed

	if err := tx.Raw("SELECT SUM(collected) as collected_money FROM posts WHERE date > ?", date).Scan(&stats.Collectedmoney).Error; err != nil {
		return nil, err
	}

	if err := tx.Raw("SELECT COUNT(*) as posts FROM posts WHERE date > ? AND (status = 'approved' OR status = 'expired')", date).Scan(&stats.Posts).Error; err != nil {
		return nil, err
	}

	if err := tx.Raw("SELECT SUM(likes) as likes FROM posts WHERE date > ?", date).Scan(&stats.Likes).Error; err != nil {
		return nil, err
	}

	if err := tx.Raw("SELECT COUNT(*) as campaigns FROM posts WHERE date > ? AND status = 'approved'", date).Scan(&stats.Activecampaigns).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

// func (s *Server) CategoryStats(ctx context.Context, req *pb.CategoryStatsRequest) (*pb.CategoryStatsResponse, error) {

// 	log.Println("Campaign Request List started")
// 	log.Println("Data collected", req)
// 	var page, limit int64
// 	page, limit = int64(req.Page), int64(req.Limit)
// 	// pagination purpose -
// 	if req.Page <= 0 {
// 		page = 1
// 	}
// 	if req.Limit <= 0 {
// 		limit = 10
// 	}
// 	offset := (page - 1) * limit
// 	var posts []*pb.Post
// 	sqlQuery := "SELECT * FROM posts"
// 	sqlQuery += " ORDER BY views DESC,likes DESC LIMIT ? OFFSET ?"

// 	if err := s.H.DB.Raw(sqlQuery, limit, offset).Scan(&posts).Error; err != nil {
// 		return &pb.PostStatsResponse{
// 			Status:   http.StatusBadRequest,
// 			Response: "couldn't get posts from DB",
// 		}, err
// 	}
// 	log.Println("Data Recieved: ", req)
// 	return &pb.CategoryStatsResponse{
// 		Status:   http.StatusOK,
// 		Response: "",
// 	}, nil

// }
