#include "opencv/cv.h"
//#include <opencv2/core/core.hpp>        // Basic OpenCV structures 
#include "opencv2/highgui/highgui.hpp"
#include "opencv2/imgproc/imgproc.hpp"
//#inlcude <omp.h>
#include<iostream>
using namespace cv;
using namespace std;
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
	double phi3;
	double sita3;
	double gama3;
	Mat ImgIndex;
	Mat ImgIndex1;
	Mat ImgIndex2;
	Mat ImgIndex22;
	Mat ImgIndex3;
	Mat ImgIndex4;
	Mat myTest;        //to  do test,by xiefanfan
	Mat orgImg;
	Mat dstImg;
	Mat dstImage1;
	Mat dstImage2;
	Mat dstImage3;
	Mat dstImage4;
	CFishEyeCorrect();
	~CFishEyeCorrect();
	Rect correctArea;
	void ParamFix();
	void CalculateParam(int width, int height);
	void Fix(int &u,int &v,double &x,double &y,const double &phi,const double &sita,const double &gama);
	void FishImgIndex(const Mat& ormImg);
	//void ImageCorrect(const Mat& orgImg,Mat& dstImg1,Mat& dstImg2,Mat& dstImg3);
	static void * pth_fun1(void *pth);
	static void * pth_fun2(void *pth);
	static void * pth_fun3(void *pth);
	static void * pth_fun4(void *pth);
	void ImageCorrect( );
	Rect GetArea(const Mat &inputImage);

	int lineWidth ;
	int eleWidth;
	int orgImgLineWidht;
	int orgImgEleWidht;
	int eleSize;


private:
	void FindCorrectArea(const Mat &inputImage);
	const double piParam;
};
