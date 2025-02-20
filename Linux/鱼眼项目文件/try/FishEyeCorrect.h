#include "opencv/cv.h"
//#include <opencv2/core/core.hpp>        // Basic OpenCV structures 
#include "opencv2/highgui/highgui.hpp"
#inlcude <omp.h>

using namespace cv;

class CFishEyeCorrect
{
public:
	int w;
	int h;
	double cx;
	double cy;
	double f;
	double r;
	double phi0;
	double sita0;
	double gama0;
	double phi1;
	double sita1;
	double gama1;
	double phi2;
	double sita2;
	double gama2;
	Mat ImgIndex1;
	Mat ImgIndex2;
	Mat ImgIndex3;
	CFishEyeCorrect();
	~CFishEyeCorrect();
	Rect correctArea;
	void ParamFix();
	void CalculateParam(int width, int height);
	void Fix(int &u,int &v,double &x,double &y,const double &phi,const double &sita,const double &gama);
	void FishImgIndex(const Mat& ormImg);
//	void ImageCorrect(const Mat& orgImg,Mat& dstImg1,Mat& dstImg2,Mat& dstImg3);
void ImageCorrect(const Mat& orgImg,Mat& dstImg,Mat& ImgIndex );
	Rect GetArea(const Mat &inputImage);
private:
	void FindCorrectArea(const Mat &inputImage);
	const double piParam;
};
