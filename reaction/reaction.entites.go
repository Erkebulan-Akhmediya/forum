package reaction

type reaction struct {
	id, userId int
	reaction   string
}

type postReaction struct {
	reaction
	postId int
}

type commentReaction struct {
	reaction
	commentId int
}
