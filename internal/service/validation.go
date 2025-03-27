package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	pb "server-data/api/validation"
)

type ValidationService struct {
	pb.UnimplementedValidationServer
}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (s *ValidationService) GetValidation(ctx context.Context, req *pb.GetValidationRequest) (*pb.GetValidationReply, error) {
	str, err := GenerateCode(int(req.Length), req.Type)
	if err != nil {
		return &pb.GetValidationReply{}, err
	}

	return &pb.GetValidationReply{
		Code: str,
	}, nil
}

// GenerateCode 根据长度和类型生成验证码
// length: 验证码长度
// codeType: 验证码类型 (DEFAULT, DIGIT, LETTER, MIXED)
func GenerateCode(length int, codeType pb.TYPE) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive")
	}

	// 定义字符集
	var chars string
	switch codeType {
	case 1:
		chars = "0123456789"
	case 2:
		chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 3:
		chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=[]{}|;:,.<>?"
	default: // DEFAULT
		chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	var sb strings.Builder
	charLen := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, charLen)
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}
		sb.WriteByte(chars[n.Int64()])
	}

	return sb.String(), nil
}
